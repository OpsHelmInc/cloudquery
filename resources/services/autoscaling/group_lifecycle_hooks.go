package autoscaling

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/autoscaling/models"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func groupLifecycleHooks() *schema.Table {
	tableName := "aws_autoscaling_group_lifecycle_hooks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html`,
		Resolver:    fetchAutoscalingGroupLifecycleHooks,
		Transform:   transformers.TransformWithStruct(&types.LifecycleHook{}, transformers.WithPrimaryKeyComponents("LifecycleHookName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "group_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAutoscalingGroupLifecycleHooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(models.AutoScalingGroupWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAutoscaling).Autoscaling
	config := autoscaling.DescribeLifecycleHooksInput{AutoScalingGroupName: p.AutoScalingGroupName}

	output, err := svc.DescribeLifecycleHooks(ctx, &config, func(options *autoscaling.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if isAutoScalingGroupNotExistsError(err) {
			return nil
		}
		return err
	}
	res <- output.LifecycleHooks
	return nil
}
