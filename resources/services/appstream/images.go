// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_images",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html`,
		Resolver:    fetchAppstreamImages,
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
				Name:     "applications",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Applications"),
			},
			{
				Name:     "appstream_agent_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppstreamAgentVersion"),
			},
			{
				Name:     "base_image_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BaseImageArn"),
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
				Name:     "dynamic_app_providers_enabled",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DynamicAppProvidersEnabled"),
			},
			{
				Name:     "image_builder_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageBuilderName"),
			},
			{
				Name:     "image_builder_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ImageBuilderSupported"),
			},
			{
				Name:     "image_errors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageErrors"),
			},
			{
				Name:     "image_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImagePermissions"),
			},
			{
				Name:     "image_shared_with_others",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageSharedWithOthers"),
			},
			{
				Name:     "latest_appstream_agent_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestAppstreamAgentVersion"),
			},
			{
				Name:     "platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "public_base_image_released_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("PublicBaseImageReleasedDate"),
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
				Name:     "supported_instance_families",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedInstanceFamilies"),
			},
			{
				Name:     "visibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Visibility"),
			},
		},
	}
}
