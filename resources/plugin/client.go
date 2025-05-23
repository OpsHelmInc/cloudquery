package plugin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/spec"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/scheduler"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/state"
)

const maxMsgSize = 100 * 1024 * 1024 // 100 MiB

type Client struct {
	plugin.UnimplementedDestination
	scheduler *scheduler.Scheduler
	client    schema.ClientMeta
	logger    zerolog.Logger
	options   plugin.NewClientOptions
	allTables schema.Tables
}

func New(ctx context.Context, logger zerolog.Logger, specBytes []byte, options plugin.NewClientOptions, overrideConfig *aws.Config) (plugin.Client, error) {
	var s spec.Spec
	c := &Client{
		options:   options,
		logger:    logger,
		allTables: GetTables(),
	}
	if options.NoConnection {
		return c, nil
	}
	var err error
	if err := json.Unmarshal(specBytes, &s); err != nil {
		return nil, err
	}
	s.SetDefaults()
	if err := s.Validate(); err != nil {
		return nil, err
	}
	c.client, err = client.Configure(ctx, logger, s, overrideConfig)
	if err != nil {
		return nil, err
	}

	c.scheduler = scheduler.NewScheduler(
		scheduler.WithConcurrency(s.Concurrency),
		scheduler.WithLogger(logger),
		scheduler.WithStrategy(*s.Scheduler),
	)
	return c, nil
}

func (*Client) Close(_ context.Context) error {
	return nil
}

func (c *Client) Tables(_ context.Context, options plugin.TableOptions) (schema.Tables, error) {
	return c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
}

func (c *Client) Sync(ctx context.Context, options plugin.SyncOptions, res chan<- message.SyncMessage, resources chan<- *schema.Resource) error {
	if c.options.NoConnection {
		return fmt.Errorf("no connection")
	}
	tt, err := c.allTables.FilterDfs(options.Tables, options.SkipTables, options.SkipDependentTables)
	if err != nil {
		return err
	}

	var stateClient state.Client
	if options.BackendOptions == nil {
		c.logger.Info().Msg("No backend options provided, using no state backend")
		stateClient = &state.NoOpClient{}
	} else {
		conn, err := grpc.NewClient(options.BackendOptions.Connection,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(maxMsgSize),
				grpc.MaxCallSendMsgSize(maxMsgSize),
			),
		)
		if err != nil {
			return fmt.Errorf("failed to dial grpc source plugin at %s: %w", options.BackendOptions.Connection, err)
		}
		stateClient, err = state.NewClient(ctx, conn, options.BackendOptions.TableName)
		if err != nil {
			return fmt.Errorf("failed to create state client: %w", err)
		}
		c.logger.Info().Str("table_name", options.BackendOptions.TableName).Msg("Connected to state backend")
	}
	awsClient := c.client.(*client.Client)
	// for each sync we want to create a copy of the client, so they won't share state
	awsClient = awsClient.Duplicate()
	awsClient.SetStateClient(stateClient)
	err = c.scheduler.Sync(ctx, awsClient, tt, res, resources, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
	if err != nil {
		return err
	}
	return stateClient.Flush(ctx)
}
