package client

import (
	"context"
	"os"
	"slices"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"

	"github.com/OpsHelmInc/cloudquery/v2/client/spec"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/scheduler"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

type TestOptions struct {
	Region string

	SkipEmptyCheckColumns map[string][]string
}

func AwsMockTestHelper(t *testing.T, parentTable *schema.Table, builder func(*testing.T, *gomock.Controller) Services, testOpts TestOptions) {
	parentTable.IgnoreInTests = false
	if testOpts.Region == "" {
		testOpts.Region = "us-east-1"
	}
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	var awsSpec spec.Spec
	awsSpec.SetDefaults()
	awsSpec.UsePaidAPIs = true
	c := NewAwsClient(l, &awsSpec)
	services := builder(t, ctrl)
	services.Regions = []string{testOpts.Region}
	services.AWSConfig.Region = testOpts.Region
	c.accountMutex["testAccount"] = &sync.Mutex{}
	c.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
	c.Partition = "aws"
	tables := schema.Tables{parentTable}

	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	validateTagStructure(t, tables)
	validateMultiplexers(t, parentTable)
	validateSkippedColumns(t, tables)
	sc := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), &c, tables)
	if err != nil {
		t.Fatal(err)
	}

	validateNoEmptyColumnsExcept(t, tables, messages, testOpts.SkipEmptyCheckColumns)
	validateTagData(t, tables, messages)
}

func AwsCreateMockClient(t *testing.T, ctrl *gomock.Controller, builder func(*testing.T, *gomock.Controller) Services, testOpts TestOptions) Client {
	if testOpts.Region == "" {
		testOpts.Region = "us-east-1"
	}

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	var awsSpec spec.Spec
	awsSpec.SetDefaults()
	awsSpec.UsePaidAPIs = true
	c := NewAwsClient(l, &awsSpec)
	if builder != nil {
		services := builder(t, ctrl)
		services.Regions = []string{testOpts.Region}
		c.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
	}

	c.accountMutex["testAccount"] = &sync.Mutex{}

	c.Partition = "aws"
	return c
}

func validateTagStructure(t *testing.T, tables schema.Tables) {
	for _, table := range tables.FlattenTables() {
		t.Run(table.Name+"/validate tags", func(t *testing.T) {
			for _, column := range table.Columns {
				if column.Name != "tags" {
					continue
				}
				if column.Type != sdkTypes.ExtensionTypes.JSON {
					t.Fatalf("tags column in %s should be of type JSON", table.Name)
				}
				// TODO: Get actual field value and ensure it is of type map[string]string
			}
		})
	}
}

func validateMultiplexers(t *testing.T, parentTable *schema.Table) {
	tables := schema.Tables{parentTable}
	for _, table := range tables.FlattenTables() {
		if table.Name == parentTable.Name {
			continue
		}
		if table.Multiplex == nil {
			continue
		}
		t.Fatalf("table %s is a relation and should not have multiplexer", table.Name)
	}
}

func validateSkippedColumns(t *testing.T, tables schema.Tables) {
	for _, table := range tables.FlattenTables() {
		t.Run(table.Name+"/validate skipped columns", func(t *testing.T) {
			for _, columnName := range []string{"result_metadata"} {
				col := table.Columns.Get(columnName)
				if !ignoreNonSkippedColumns(table.Name, columnName) && col != nil {
					t.Fatalf("column %s in table %s should be skipped", columnName, table.Name)
				}
			}
		})
	}
}

func ignoreNonSkippedColumns(tableName, column string) bool {
	tableColumnNamesToIgnore := map[string]bool{}
	_, ok := tableColumnNamesToIgnore[tableName+"."+column]
	return ok
}

func validateNoEmptyColumnsExcept(t *testing.T, tables schema.Tables, messages message.SyncMessages, except map[string][]string) {
	// same as SDK's plugin.ValidateNoEmptyColumns, with exceptions for some columns
	for _, table := range tables.FlattenTables() {
		records := messages.GetInserts().GetRecordsForTable(table)
		emptyColumns := schema.FindEmptyColumns(table, records)

		emptyColumns = slices.DeleteFunc(emptyColumns, func(a string) bool {
			return slices.Contains(except[table.Name], a)
		})

		if len(emptyColumns) > 0 {
			t.Fatalf("found empty column(s): %v in %s", emptyColumns, table.Name)
		}
	}
}

func validateTagData(t *testing.T, tables schema.Tables, _ message.SyncMessages) {
	tablesWithIssues := []string{}
	for _, table := range tables.FlattenTables() {
		index := table.Columns.Index("tags")
		if index == -1 {
			continue
		}
		/*
			records := messages.GetInserts().GetRecordsForTable(table)
			for _, resource := range records {
				arr := resource.Column(index)
				for i := 0; i < arr.Len(); i++ {
					val := arr.GetOneForMarshal(i).(json.RawMessage)
						if strings.HasPrefix(string(val), "[") && strings.HasSuffix(string(val), "]") {
							tablesWithIssues = append(tablesWithIssues, table.Name)
							break
						}
				}
			}
		*/
	}
	if len(tablesWithIssues) > 0 {
		t.Fatalf("found improperly structured tags in %v", tablesWithIssues)
	}
}
