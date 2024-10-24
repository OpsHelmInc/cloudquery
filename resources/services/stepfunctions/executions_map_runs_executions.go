package stepfunctions

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func mapRunExecutions() *schema.Table {
	tableName := "aws_stepfunctions_map_run_executions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html`,
		Resolver:            fetchStepfunctionsMapRunExecutions,
		PreResourceResolver: getExecution,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeExecutionOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ExecutionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "map_run_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "state_machine_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("state_machine_arn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchStepfunctionsMapRunExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSfn).Sfn
	config := sfn.ListExecutionsInput{
		MaxResults: 1000,
		MapRunArn:  parent.Item.(*sfn.DescribeMapRunOutput).MapRunArn,
	}
	paginator := sfn.NewListExecutionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *sfn.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Executions
	}
	return nil
}
