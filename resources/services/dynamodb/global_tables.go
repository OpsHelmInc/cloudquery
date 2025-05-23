package dynamodb

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func GlobalTables() *schema.Table {
	tableName := "aws_dynamodb_global_tables"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GlobalTableDescription.html
This table only contains version 2017.11.29 (Legacy) Global Tables. See aws_dynamodb_tables for version 2019.11.21 (Current) Global Tables.
The column "tags" is always empty because global tables do not support tags. The column will be removed in a future version.
OH: The tag has been removed.
`,
		Resolver:            fetchGlobalTables,
		PreResourceResolver: getGlobalTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.GlobalTableDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("GlobalTableArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchGlobalTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	config := dynamodb.ListGlobalTablesInput{
		RegionName: aws.String(cl.Region),
	}
	// No paginator available
	for {
		output, err := svc.ListGlobalTables(ctx, &config, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.GlobalTables

		if aws.ToString(output.LastEvaluatedGlobalTableName) == "" {
			break
		}
		config.ExclusiveStartGlobalTableName = output.LastEvaluatedGlobalTableName
	}

	return nil
}

func getGlobalTable(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	table := resource.Item.(types.GlobalTable)

	response, err := svc.DescribeGlobalTable(ctx, &dynamodb.DescribeGlobalTableInput{GlobalTableName: table.GlobalTableName}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response.GlobalTableDescription
	return nil
}
