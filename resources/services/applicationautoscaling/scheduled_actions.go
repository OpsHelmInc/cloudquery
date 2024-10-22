package applicationautoscaling

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func ScheduledActions() *schema.Table {
	tableName := "aws_applicationautoscaling_scheduled_actions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScheduledAction.html`,
		Resolver:    fetchScheduledActions,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScheduledAction{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ScheduledActionARN"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchScheduledActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApplicationautoscaling).Applicationautoscaling

	config := applicationautoscaling.DescribeScheduledActionsInput{
		ServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),
	}
	paginator := applicationautoscaling.NewDescribeScheduledActionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *applicationautoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ScheduledActions
	}

	return nil
}
