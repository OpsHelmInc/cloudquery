package stepfunctions

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func mapRuns() *schema.Table {
	tableName := "aws_stepfunctions_map_runs"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeMapRun.html`,
		Resolver:            fetchStepfunctionsMapRuns,
		PreResourceResolver: getMapRun,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeMapRunOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("MapRunArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "state_machine_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			mapRunExecutions(),
		},
	}
}

func fetchStepfunctionsMapRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSfn).Sfn
	config := sfn.ListMapRunsInput{
		MaxResults:   1000,
		ExecutionArn: parent.Item.(*sfn.DescribeExecutionOutput).ExecutionArn,
	}
	paginator := sfn.NewListMapRunsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *sfn.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.MapRuns
	}
	return nil
}

func getMapRun(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSfn).Sfn
	config := sfn.DescribeMapRunInput{
		MapRunArn: resource.Item.(types.MapRunListItem).MapRunArn,
	}
	output, err := svc.DescribeMapRun(ctx, &config, func(o *sfn.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
