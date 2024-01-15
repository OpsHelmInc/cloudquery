// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_applications",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Application.html`,
		Resolver:    fetchAppstreamApplications,
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
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::AppStream::Application"),
			},
			{
				Name:     "app_block_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AppBlockArn"),
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
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "icon_s3_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IconS3Location"),
			},
			{
				Name:     "icon_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IconURL"),
			},
			{
				Name:     "instance_families",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InstanceFamilies"),
			},
			{
				Name:     "launch_parameters",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchParameters"),
			},
			{
				Name:     "launch_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchPath"),
			},
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "platforms",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Platforms"),
			},
			{
				Name:     "working_directory",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkingDirectory"),
			},
		},

		Relations: []*schema.Table{
			ApplicationFleetAssociations(),
		},
	}
}
