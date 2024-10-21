package batchwriter

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
)

type IgnoreMigrateTables struct{}

func (IgnoreMigrateTables) MigrateTables(context.Context, message.WriteMigrateTables) error {
	return nil
}

type UnimplementedDeleteStale struct{}

func (UnimplementedDeleteStale) DeleteStale(context.Context, message.WriteDeleteStales) error {
	return fmt.Errorf("DeleteStale: %w", plugin.ErrNotImplemented)
}

type UnimplementedDeleteRecord struct{}

func (UnimplementedDeleteRecord) DeleteRecord(context.Context, message.WriteDeleteRecords) error {
	return fmt.Errorf("DeleteRecord: %w", plugin.ErrNotImplemented)
}
