package dynamodb

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Exports() *schema.Table {
	tableName := "aws_dynamodb_exports"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ExportDescription.html`,
		Resolver:            listExports,
		PreResourceResolver: getExport,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "dynamodb"),
		Transform:           transformers.TransformWithStruct(&types.ExportDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ExportArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func listExports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	paginator := dynamodb.NewListExportsPaginator(svc, &dynamodb.ListExportsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *dynamodb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ExportSummaries
	}

	return nil
}

func getExport(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDynamodb).Dynamodb

	exportSummary := resource.Item.(types.ExportSummary)

	response, err := svc.DescribeExport(ctx, &dynamodb.DescribeExportInput{ExportArn: exportSummary.ExportArn}, func(options *dynamodb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = response.ExportDescription
	return nil
}
