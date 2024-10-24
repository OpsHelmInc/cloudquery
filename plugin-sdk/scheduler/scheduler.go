package scheduler

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog"
	"github.com/thoas/go-funk"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/semaphore"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/caser"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

const (
	DefaultSingleResourceMaxConcurrency    = 5
	DefaultSingleNestedTableMaxConcurrency = 5
	DefaultConcurrency                     = 50000
	DefaultMaxDepth                        = 4
	minTableConcurrency                    = 1
	minResourceConcurrency                 = 100
	otelName                               = "io.cloudquery"
)

var ErrNoTables = errors.New("no tables specified for syncing, review `tables` and `skip_tables` in your config and specify at least one table to sync")

const (
	StrategyDFS Strategy = iota
	StrategyRoundRobin
	StrategyShuffle
)

type Option func(*Scheduler)

func WithLogger(logger zerolog.Logger) Option {
	return func(s *Scheduler) {
		s.logger = logger
	}
}

func WithConcurrency(concurrency int) Option {
	return func(s *Scheduler) {
		s.concurrency = concurrency
	}
}

func WithMaxDepth(maxDepth uint64) Option {
	return func(s *Scheduler) {
		s.maxDepth = maxDepth
	}
}

func WithStrategy(strategy Strategy) Option {
	return func(s *Scheduler) {
		s.strategy = strategy
	}
}

func WithSingleNestedTableMaxConcurrency(concurrency int64) Option {
	return func(s *Scheduler) {
		s.singleNestedTableMaxConcurrency = concurrency
	}
}

func WithSingleResourceMaxConcurrency(concurrency int64) Option {
	return func(s *Scheduler) {
		s.singleResourceMaxConcurrency = concurrency
	}
}

type SyncOption func(*syncClient)

func WithSyncDeterministicCQID(deterministicCQID bool) SyncOption {
	return func(s *syncClient) {
		s.deterministicCQID = deterministicCQID
	}
}

func WithInvocationID(invocationID string) Option {
	return func(s *Scheduler) {
		s.invocationID = invocationID
	}
}

type Client interface {
	ID() string
}

type Scheduler struct {
	caser    *caser.Caser
	strategy Strategy
	maxDepth uint64
	// resourceSem is a semaphore that limits the number of concurrent resources being fetched
	resourceSem *semaphore.Weighted
	// tableSem is a semaphore that limits the number of concurrent tables being fetched
	tableSems []*semaphore.Weighted
	// Logger to call, this logger is passed to the serve.Serve Client, if not defined Serve will create one instead.
	logger      zerolog.Logger
	concurrency int
	// This Map holds all of the concurrency semaphores for each table+client pair.
	singleTableConcurrency sync.Map
	// The maximum number of go routines that can be spawned for a single table+client pair
	singleNestedTableMaxConcurrency int64

	// The maximum number of go routines that can be spawned for a specific resource
	singleResourceMaxConcurrency int64

	// Controls how records are constructed on the source side.
	batchSettings *BatchSettings

	invocationID string
}

type syncClient struct {
	tables            schema.Tables
	client            schema.ClientMeta
	scheduler         *Scheduler
	deterministicCQID bool
	// status sync metrics
	metrics      *Metrics
	logger       zerolog.Logger
	invocationID string
}

func NewScheduler(opts ...Option) *Scheduler {
	s := Scheduler{
		caser:                           caser.New(),
		concurrency:                     DefaultConcurrency,
		maxDepth:                        DefaultMaxDepth,
		singleResourceMaxConcurrency:    DefaultSingleResourceMaxConcurrency,
		singleNestedTableMaxConcurrency: DefaultSingleNestedTableMaxConcurrency,
		batchSettings: &BatchSettings{
			MaxRows: DefaultBatchMaxRows,
			Timeout: DefaultBatchTimeout,
		},
	}
	for _, opt := range opts {
		opt(&s)
	}
	// This is very similar to the concurrent web crawler problem with some minor changes.
	// We are using DFS/Round-Robin to make sure memory usage is capped at O(h) where h is the height of the tree.
	tableConcurrency := max(s.concurrency/minResourceConcurrency, minTableConcurrency)
	resourceConcurrency := tableConcurrency * minResourceConcurrency
	s.tableSems = make([]*semaphore.Weighted, s.maxDepth)
	for i := uint64(0); i < s.maxDepth; i++ {
		s.tableSems[i] = semaphore.NewWeighted(int64(tableConcurrency))
		// reduce table concurrency logarithmically for every depth level
		tableConcurrency = max(tableConcurrency/2, minTableConcurrency)
	}
	s.resourceSem = semaphore.NewWeighted(int64(resourceConcurrency))

	return &s
}

// SyncAll is mostly used for testing as it will sync all tables and can run out of memory
// in the real world. Should use Sync for production.
func (s *Scheduler) SyncAll(ctx context.Context, client schema.ClientMeta, tables schema.Tables) (message.SyncMessages, error) {
	res := make(chan message.SyncMessage)
	resourceCh := make(chan *schema.Resource)
	var err error
	go func() {
		defer close(res)
		defer close(resourceCh)
		err = s.Sync(ctx, client, tables, res, resourceCh)
	}()
	// nolint:prealloc
	var messages message.SyncMessages
	for msg := range res {
		messages = append(messages, msg)
	}
	return messages, err
}

func (s *Scheduler) Sync(ctx context.Context, client schema.ClientMeta, tables schema.Tables, res chan<- message.SyncMessage, resourceCh chan<- *schema.Resource, opts ...SyncOption) error {
	ctx, span := otel.Tracer(otelName).Start(ctx,
		"sync",
		trace.WithAttributes(attribute.Key("sync.invocation.id").String(s.invocationID)),
	)
	defer span.End()
	if len(tables) == 0 {
		return ErrNoTables
	}

	syncClient := &syncClient{
		metrics:      &Metrics{TableClient: make(map[string]map[string]*TableClientMetrics)},
		tables:       tables,
		client:       client,
		scheduler:    s,
		logger:       s.logger,
		invocationID: s.invocationID,
	}
	for _, opt := range opts {
		opt(syncClient)
	}

	if maxDepth(tables) > s.maxDepth {
		return fmt.Errorf("max depth exceeded, max depth is %d", s.maxDepth)
	}

	// send migrate messages first
	for _, tableOriginal := range tables.FlattenTables() {
		migrateMessage := &message.SyncMigrateTable{
			Table: tableOriginal.Copy(tableOriginal.Parent),
		}
		if syncClient.deterministicCQID {
			schema.CqIDAsPK(migrateMessage.Table)
		}
		res <- migrateMessage
	}

	resources := make(chan *schema.Resource)
	go func() {
		defer close(resources)
		switch s.strategy {
		case StrategyDFS:
			syncClient.syncDfs(ctx, resources)
		case StrategyRoundRobin:
			syncClient.syncRoundRobin(ctx, resources)
		case StrategyShuffle:
			syncClient.syncShuffle(ctx, resources)
		default:
			panic(fmt.Errorf("unknown scheduler %s", s.strategy.String()))
		}
	}()

	b := s.batchSettings.getBatcher(ctx, res, s.logger)
	defer b.close()    // wait for all resources to be processed
	done := ctx.Done() // no need to do the lookups in loop
	for resource := range resources {
		select {
		case <-done:
			s.logger.Debug().Msg("sync context cancelled")
			return context.Cause(ctx)
		default:
			b.process(resource)
		}

		select {
		case resourceCh <- resource:
		default:
			s.logger.Error().Msg("failed to publish resource to channel")
		}
	}

	return context.Cause(ctx)
}

func (s *syncClient) logTablesMetrics(tables schema.Tables, client Client) {
	clientName := client.ID()
	for _, table := range tables {
		metrics := s.metrics.TableClient[table.Name][clientName]
		duration := metrics.Duration.Load()
		if duration == nil {
			// This can happen for a relation when there are no resources to resolve from the parent
			duration = new(time.Duration)
		}
		s.logger.Debug().Str("table", table.Name).Str("client", clientName).Uint64("resources", metrics.Resources).Dur("duration_ms", *duration).Uint64("errors", metrics.Errors).Msg("table sync finished")
		s.logTablesMetrics(table.Relations, client)
	}
}

func (s *syncClient) resolveResource(ctx context.Context, table *schema.Table, client schema.ClientMeta, parent *schema.Resource, item any) *schema.Resource {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()
	resource := schema.NewResourceData(table, parent, item)
	objectStartTime := time.Now()
	clientID := client.ID()
	tableMetrics := s.metrics.TableClient[table.Name][clientID]
	logger := s.logger.With().Str("table", table.Name).Str("client", clientID).Logger()
	defer func() {
		if err := recover(); err != nil {
			stack := fmt.Sprintf("%s\n%s", err, string(debug.Stack()))
			logger.Error().Interface("error", err).TimeDiff("duration", time.Now(), objectStartTime).Str("stack", stack).Msg("resource resolver finished with panic")
			atomic.AddUint64(&tableMetrics.Panics, 1)
		}
	}()
	if table.PreResourceResolver != nil {
		if err := table.PreResourceResolver(ctx, client, resource); err != nil {
			logger.Error().Err(err).Msg("pre resource resolver failed")
			atomic.AddUint64(&tableMetrics.Errors, 1)
			return nil
		}
	}

	for _, c := range table.Columns {
		s.resolveColumn(ctx, logger, tableMetrics, client, resource, c)
	}

	if table.PostResourceResolver != nil {
		if err := table.PostResourceResolver(ctx, client, resource); err != nil {
			logger.Error().Stack().Err(err).Msg("post resource resolver finished with error")
			atomic.AddUint64(&tableMetrics.Errors, 1)
		}
	}

	tableMetrics.OtelResourcesAdd(ctx, 1)
	atomic.AddUint64(&tableMetrics.Resources, 1)
	return resource
}

func (s *syncClient) resolveColumn(ctx context.Context, logger zerolog.Logger, tableMetrics *TableClientMetrics, client schema.ClientMeta, resource *schema.Resource, c schema.Column) {
	columnStartTime := time.Now()
	defer func() {
		if err := recover(); err != nil {
			stack := fmt.Sprintf("%s\n%s", err, string(debug.Stack()))
			logger.Error().Str("column", c.Name).Interface("error", err).TimeDiff("duration", time.Now(), columnStartTime).Str("stack", stack).Msg("column resolver finished with panic")
			atomic.AddUint64(&tableMetrics.Panics, 1)
		}
	}()

	if c.Resolver != nil {
		if err := c.Resolver(ctx, client, resource, c); err != nil {
			logger.Error().Err(err).Msg("column resolver finished with error")
			atomic.AddUint64(&tableMetrics.Errors, 1)
		}
	} else {
		// base use case: try to get column with CamelCase name
		v := funk.Get(resource.GetItem(), s.scheduler.caser.ToPascal(c.Name), funk.WithAllowZero())
		if v != nil {
			err := resource.Set(c.Name, v)
			if err != nil {
				logger.Error().Err(err).Msg("column resolver finished with error")
				atomic.AddUint64(&tableMetrics.Errors, 1)
			}
		}
	}
}

func maxDepth(tables schema.Tables) uint64 {
	var depth uint64
	if len(tables) == 0 {
		return 0
	}
	for _, table := range tables {
		newDepth := 1 + maxDepth(table.Relations)
		if newDepth > depth {
			depth = newDepth
		}
	}
	return depth
}
