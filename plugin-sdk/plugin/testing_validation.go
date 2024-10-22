package plugin

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func ValidateNoEmptyColumns(t *testing.T, tables schema.Tables, messages message.SyncMessages) {
	for _, table := range tables.FlattenTables() {
		records := messages.GetInserts().GetRecordsForTable(table)
		emptyColumns := schema.FindEmptyColumns(table, records)
		if len(emptyColumns) > 0 {
			t.Fatalf("found empty column(s): %v in %s", emptyColumns, table.Name)
		}
	}
}
