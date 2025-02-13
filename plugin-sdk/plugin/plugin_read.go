package plugin

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

// readAll is used in tests to read all records from a table.
func (p *Plugin) readAll(ctx context.Context, table *schema.Table) ([]arrow.Record, error) {
	var err error
	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = p.client.Read(ctx, table, ch)
	}()
	// nolint:prealloc
	var records []arrow.Record
	for record := range ch {
		records = append(records, sliceToSingleRowRecord(record)...)
	}

	return records, err
}

func sliceToSingleRowRecord(record arrow.Record) []arrow.Record {
	result := make([]arrow.Record, record.NumRows())
	for i := int64(0); i < record.NumRows(); i++ {
		result[i] = record.NewSlice(i, i+1)
	}
	return result
}
