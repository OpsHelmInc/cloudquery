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

func ScalableTargets() *schema.Table {
	tableName := "aws_applicationautoscaling_scalable_targets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalableTarget.html`,
		Resolver:    fetchScalableTargets,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalableTarget{}, transformers.WithPrimaryKeyComponents("ResourceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("ScalableTargetARN"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchScalableTargets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApplicationautoscaling).Applicationautoscaling

	config := applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: types.ServiceNamespace(cl.AutoscalingNamespace),
	}
	paginator := applicationautoscaling.NewDescribeScalableTargetsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *applicationautoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ScalableTargets
	}
	return nil
}
