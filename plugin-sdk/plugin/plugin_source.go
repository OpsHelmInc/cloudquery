package plugin

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/glob"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

type BackendOptions struct {
	TableName  string
	Connection string
}

type SyncOptions struct {
	Tables              []string
	SkipTables          []string
	SkipDependentTables bool
	DeterministicCQID   bool
	BackendOptions      *BackendOptions
}

type SourceClient interface {
	Close(ctx context.Context) error
	Tables(ctx context.Context, options TableOptions) (schema.Tables, error)
	Sync(ctx context.Context, options SyncOptions, res chan<- message.SyncMessage, resourceCh chan<- *schema.Resource) error
}

func MatchesTable(name string, includeTablesPattern []string, skipTablesPattern []string) bool {
	for _, pattern := range skipTablesPattern {
		if glob.Glob(pattern, name) {
			return false
		}
	}
	for _, pattern := range includeTablesPattern {
		if glob.Glob(pattern, name) {
			return true
		}
	}
	return false
}

type NewSourceClientFunc func(context.Context, zerolog.Logger, any) (SourceClient, error)

// NewSourcePlugin returns a new CloudQuery Plugin with the given name, version and implementation.
// Source plugins only support read operations. For Read & Write plugin use NewPlugin.
func NewSourcePlugin(name string, version string, newClient NewSourceClientFunc, options ...Option) *Plugin {
	newClientWrapper := func(ctx context.Context, logger zerolog.Logger, spec []byte, _ NewClientOptions, _ *aws.Config) (Client, error) {
		sourceClient, err := newClient(ctx, logger, spec)
		if err != nil {
			return nil, err
		}
		wrapperClient := struct {
			SourceClient
			UnimplementedDestination
		}{
			SourceClient: sourceClient,
		}
		return wrapperClient, nil
	}
	return NewPlugin(name, version, newClientWrapper, options...)
}

func (p *Plugin) SyncAll(ctx context.Context, options SyncOptions) (message.SyncMessages, error) {
	var err error
	ch := make(chan message.SyncMessage)
	resourceCh := make(chan *schema.Resource)
	go func() {
		defer close(ch)
		err = p.Sync(ctx, options, ch, resourceCh)
	}()
	// nolint:prealloc
	var resources message.SyncMessages
	for resource := range ch {
		resources = append(resources, resource)
	}
	return resources, err
}

// Sync is syncing data from the requested tables in spec to the given channel
func (p *Plugin) Sync(ctx context.Context, options SyncOptions, res chan<- message.SyncMessage, resourceCh chan<- *schema.Resource) error {
	if !p.mu.TryLock() {
		return fmt.Errorf("plugin already in use")
	}
	defer p.mu.Unlock()
	if p.client == nil {
		return fmt.Errorf("plugin not initialized. call Init() first")
	}

	if err := p.client.Sync(ctx, options, res, resourceCh); err != nil {
		return fmt.Errorf("failed to sync unmanaged client: %w", err)
	}

	return nil
}
