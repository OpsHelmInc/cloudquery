package mixedbatchwriter

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
)

type IgnoreMigrateTableBatch struct{}

func (IgnoreMigrateTableBatch) MigrateTableBatch(context.Context, message.WriteMigrateTables) error {
	return nil
}

type UnimplementedDeleteStaleBatch struct{}

func (UnimplementedDeleteStaleBatch) DeleteStaleBatch(context.Context, message.WriteDeleteStales) error {
	return fmt.Errorf("DeleteStaleBatch: %w", plugin.ErrNotImplemented)
}

type UnimplementedDeleteRecordsBatch struct{}

func (UnimplementedDeleteRecordsBatch) DeleteRecordsBatch(context.Context, message.WriteDeleteRecords) error {
	return fmt.Errorf("DeleteRecordsBatch: %w", plugin.ErrNotImplemented)
}
