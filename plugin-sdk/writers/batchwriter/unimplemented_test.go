package batchwriter_test

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/writers/batchwriter"
)

type testDummyClient struct {
	batchwriter.IgnoreMigrateTables
	batchwriter.UnimplementedDeleteStale
	batchwriter.UnimplementedDeleteRecord
}

func (testDummyClient) WriteTableBatch(context.Context, string, message.WriteInserts) error {
	return nil
}

var _ batchwriter.Client = (*testDummyClient)(nil)
