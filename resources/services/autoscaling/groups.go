// Code generated by codegen; DO NOT EDIT.

package autoscaling

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_groups",
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html`,
		Resolver:    fetchAutoscalingGroups,
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
				Name:     "load_balancers",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancers,
			},
			{
				Name:     "load_balancer_target_groups",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancerTargetGroups,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "default_cooldown",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultCooldown"),
			},
			{
				Name:     "desired_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DesiredCapacity"),
			},
			{
				Name:     "health_check_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckType"),
			},
			{
				Name:     "max_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSize"),
			},
			{
				Name:     "min_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinSize"),
			},
			{
				Name:     "capacity_rebalance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CapacityRebalance"),
			},
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Context"),
			},
			{
				Name:     "default_instance_warmup",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultInstanceWarmup"),
			},
			{
				Name:     "desired_capacity_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DesiredCapacityType"),
			},
			{
				Name:     "enabled_metrics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnabledMetrics"),
			},
			{
				Name:     "health_check_grace_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckGracePeriod"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:     "launch_configuration_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchConfigurationName"),
			},
			{
				Name:     "launch_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LaunchTemplate"),
			},
			{
				Name:     "load_balancer_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("LoadBalancerNames"),
			},
			{
				Name:     "max_instance_lifetime",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxInstanceLifetime"),
			},
			{
				Name:     "mixed_instances_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MixedInstancesPolicy"),
			},
			{
				Name:     "new_instances_protected_from_scale_in",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("NewInstancesProtectedFromScaleIn"),
			},
			{
				Name:     "placement_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlacementGroup"),
			},
			{
				Name:     "predicted_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PredictedCapacity"),
			},
			{
				Name:     "service_linked_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceLinkedRoleARN"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "suspended_processes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SuspendedProcesses"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "target_group_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TargetGroupARNs"),
			},
			{
				Name:     "termination_policies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TerminationPolicies"),
			},
			{
				Name:     "vpc_zone_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VPCZoneIdentifier"),
			},
			{
				Name:     "warm_pool_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WarmPoolConfiguration"),
			},
			{
				Name:     "warm_pool_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("WarmPoolSize"),
			},
			{
				Name:     "notification_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationConfigurations"),
			},
		},

		Relations: []*schema.Table{
			GroupScalingPolicies(),
			GroupLifecycleHooks(),
		},
	}
}
