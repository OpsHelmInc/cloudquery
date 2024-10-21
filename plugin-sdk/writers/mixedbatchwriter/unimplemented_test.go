package mixedbatchwriter_test

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/writers/mixedbatchwriter"
)

type testDummyClient struct {
	mixedbatchwriter.IgnoreMigrateTableBatch
	mixedbatchwriter.UnimplementedDeleteStaleBatch
	mixedbatchwriter.UnimplementedDeleteRecordsBatch
}

func (testDummyClient) InsertBatch(context.Context, message.WriteInserts) error {
	return nil
}

var _ mixedbatchwriter.Client = (*testDummyClient)(nil)
