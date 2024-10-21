package streamingbatchwriter_test

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/writers/streamingbatchwriter"
)

type testDummyClient struct {
	streamingbatchwriter.IgnoreMigrateTable
	streamingbatchwriter.UnimplementedDeleteStale
	streamingbatchwriter.UnimplementedDeleteRecords
}

func (testDummyClient) WriteTable(context.Context, <-chan *message.WriteInsert) error {
	return nil
}

var _ streamingbatchwriter.Client = (*testDummyClient)(nil)
