// Code generated by codegen; DO NOT EDIT.

package autoscaling

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func GroupLifecycleHooks() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_group_lifecycle_hooks",
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html`,
		Resolver:    fetchAutoscalingGroupLifecycleHooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("autoscaling"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "default_result",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultResult"),
			},
			{
				Name:     "global_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("GlobalTimeout"),
			},
			{
				Name:     "heartbeat_timeout",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HeartbeatTimeout"),
			},
			{
				Name:     "lifecycle_hook_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LifecycleHookName"),
			},
			{
				Name:     "lifecycle_transition",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LifecycleTransition"),
			},
			{
				Name:     "notification_metadata",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotificationMetadata"),
			},
			{
				Name:     "notification_target_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NotificationTargetARN"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleARN"),
			},
		},
	}
}
