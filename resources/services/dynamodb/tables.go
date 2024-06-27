package dynamodb

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Tables() *schema.Table {
	tableName := "aws_dynamodb_tables"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TableDescription.html`,
		Resolver:            fetchDynamodbTables,
		PreResourceResolver: getTable,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.TableDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDynamodbTableTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TableArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "archival_summary",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ArchivalSummary"),
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			tableReplicaAutoScalings(),
			tableContinuousBackups(),
		},
	}
}

func fetchDynamodbTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	config := dynamodb.ListTablesInput{}
	// No paginator available
	for {
		output, err := svc.ListTables(ctx, &config, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.TableNames

		if aws.ToString(output.LastEvaluatedTableName) == "" {
			break
		}
		config.ExclusiveStartTableName = output.LastEvaluatedTableName
	}

	return nil
}

func getTable(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	tableName := resource.Item.(string)

	response, err := svc.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: &tableName}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response.Table
	return nil
}

func resolveDynamodbTableTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	table := resource.Item.(*types.TableDescription)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb
	var tags []types.Tag
	input := &dynamodb.ListTagsOfResourceInput{
		ResourceArn: table.TableArn,
	}
	// No paginator available
	for {
		res, err := svc.ListTagsOfResource(ctx, input, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		tags = append(tags, res.Tags...)
		if aws.ToString(res.NextToken) == "" {
			break
		}
		input.NextToken = res.NextToken
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
