package athena

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func workGroupQueryExecutions() *schema.Table {
	tableName := "aws_athena_work_group_query_executions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html`,
		Resolver:            fetchAthenaWorkGroupQueryExecutions,
		PreResourceResolver: getWorkGroupQueryExecution,
		Transform:           transformers.TransformWithStruct(&types.QueryExecution{}, transformers.WithPrimaryKeyComponents("QueryExecutionId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "work_group_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAthenaWorkGroupQueryExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAthena).Athena
	wg := parent.Item.(types.WorkGroup)
	paginator := athena.NewListQueryExecutionsPaginator(svc, &athena.ListQueryExecutionsInput{WorkGroup: wg.Name})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *athena.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.QueryExecutionIds
	}
	return nil
}

func getWorkGroupQueryExecution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAthena).Athena

	d := resource.Item.(string)
	dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
		QueryExecutionId: aws.String(d),
	}, func(options *athena.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.QueryExecution
	return nil
}
