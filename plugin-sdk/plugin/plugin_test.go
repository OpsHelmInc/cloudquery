package plugin

import (
	"context"
	"testing"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/rs/zerolog"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

type testPluginClient struct {
	messages  message.SyncMessages
	resources schema.Resources
}

func newTestPluginClient(context.Context, zerolog.Logger, []byte, NewClientOptions, *aws.Config) (Client, error) {
	return &testPluginClient{}, nil
}

func (*testPluginClient) GetSpec() any {
	return &struct{}{}
}

func (*testPluginClient) Tables(context.Context, TableOptions) (schema.Tables, error) {
	return schema.Tables{}, nil
}

func (*testPluginClient) Read(context.Context, *schema.Table, chan<- arrow.Record) error {
	return nil
}

func (c *testPluginClient) Sync(_ context.Context, _ SyncOptions, res chan<- message.SyncMessage, resourceCh chan<- *schema.Resource) error {
	for _, msg := range c.messages {
		res <- msg
	}

	for _, r := range c.resources {
		resourceCh <- r
	}
	return nil
}

func (c *testPluginClient) Write(_ context.Context, res <-chan message.WriteMessage) error {
	for msg := range res {
		switch m := msg.(type) {
		case *message.WriteMigrateTable:
			c.messages = append(c.messages, &message.SyncMigrateTable{
				Table: m.Table,
			})
		case *message.WriteInsert:
			c.messages = append(c.messages, &message.SyncInsert{
				Record: m.Record,
			})
		default:
			panic("unknown message")
		}
	}
	return nil
}

func (*testPluginClient) Close(context.Context) error {
	return nil
}

func TestPluginSuccess(t *testing.T) {
	ctx := context.Background()
	p := NewPlugin("test", "v1.0.0", newTestPluginClient)
	if err := p.Init(ctx, []byte(""), NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	tables, err := p.Tables(ctx, TableOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(tables) != 0 {
		t.Fatal("expected 0 tables")
	}
	if err := p.WriteAll(ctx, nil); err != nil {
		t.Fatal(err)
	}
	if err := p.WriteAll(ctx, []message.WriteMessage{
		&message.WriteMigrateTable{},
	}); err != nil {
		t.Fatal(err)
	}
	if len(p.client.(*testPluginClient).messages) != 1 {
		t.Fatal("expected 1 message")
	}

	messages, err := p.SyncAll(ctx, SyncOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(messages) != 1 {
		t.Fatal("expected 1 message")
	}

	if err := p.Close(ctx); err != nil {
		t.Fatal(err)
	}
}
