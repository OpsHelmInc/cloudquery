// Code generated by codegen; DO NOT EDIT.

package ecs

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ClusterServices() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecs_cluster_services",
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Service.html`,
		Resolver:    fetchEcsClusterServices,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ecs"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::ECS::Service"),
			},
			{
				Name:     "availability_zone_rebalancing",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneRebalancing"),
			},
			{
				Name:     "capacity_provider_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CapacityProviderStrategy"),
			},
			{
				Name:     "cluster_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "deployment_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeploymentConfiguration"),
			},
			{
				Name:     "deployment_controller",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeploymentController"),
			},
			{
				Name:     "deployments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Deployments"),
			},
			{
				Name:     "desired_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DesiredCount"),
			},
			{
				Name:     "enable_ecs_managed_tags",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableECSManagedTags"),
			},
			{
				Name:     "enable_execute_command",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableExecuteCommand"),
			},
			{
				Name:     "events",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Events"),
			},
			{
				Name:     "health_check_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckGracePeriodSeconds"),
			},
			{
				Name:     "launch_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchType"),
			},
			{
				Name:     "load_balancers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoadBalancers"),
			},
			{
				Name:     "network_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkConfiguration"),
			},
			{
				Name:     "pending_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingCount"),
			},
			{
				Name:     "placement_constraints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementConstraints"),
			},
			{
				Name:     "placement_strategy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementStrategy"),
			},
			{
				Name:     "platform_family",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformFamily"),
			},
			{
				Name:     "platform_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformVersion"),
			},
			{
				Name:     "propagate_tags",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PropagateTags"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "running_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RunningCount"),
			},
			{
				Name:     "scheduling_strategy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SchedulingStrategy"),
			},
			{
				Name:     "service_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceName"),
			},
			{
				Name:     "service_registries",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServiceRegistries"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "task_definition",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskDefinition"),
			},
			{
				Name:     "task_sets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TaskSets"),
			},
		},
	}
}
