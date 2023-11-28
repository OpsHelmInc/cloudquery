// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AutoScalingConfigurations() *schema.Table {
	return &schema.Table{
		Name:                "aws_apprunner_auto_scaling_configurations",
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html`,
		Resolver:            fetchApprunnerAutoScalingConfigurations,
		PreResourceResolver: getAutoScalingConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer("apprunner"),
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
				Resolver: schema.PathResolver("AutoScalingConfigurationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("AutoScalingConfigurationArn"),
			},
			{
				Name:     "auto_scaling_configuration_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingConfigurationName"),
			},
			{
				Name:     "auto_scaling_configuration_revision",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AutoScalingConfigurationRevision"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
			},
			{
				Name:     "has_associated_service",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasAssociatedService"),
			},
			{
				Name:     "is_default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsDefault"),
			},
			{
				Name:     "latest",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Latest"),
			},
			{
				Name:     "max_concurrency",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxConcurrency"),
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
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
