// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ImageBuilders() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_image_builders",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_ImageBuilder.html`,
		Resolver:    fetchAppstreamImageBuilders,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "access_endpoints",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessEndpoints"),
			},
			{
				Name:     "appstream_agent_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppstreamAgentVersion"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "domain_join_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainJoinInfo"),
			},
			{
				Name:     "enable_default_internet_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDefaultInternetAccess"),
			},
			{
				Name:     "iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleArn"),
			},
			{
				Name:     "image_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageArn"),
			},
			{
				Name:     "image_builder_errors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageBuilderErrors"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "network_access_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkAccessConfiguration"),
			},
			{
				Name:     "platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_change_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateChangeReason"),
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
		},
	}
}
