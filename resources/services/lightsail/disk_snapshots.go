// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DiskSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_disk_snapshots",
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_DiskSnapshot.html`,
		Resolver:    fetchLightsailDiskSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "disk_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "from_disk_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromDiskArn"),
			},
			{
				Name:     "from_disk_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromDiskName"),
			},
			{
				Name:     "from_instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromInstanceArn"),
			},
			{
				Name:     "from_instance_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FromInstanceName"),
			},
			{
				Name:     "is_from_auto_snapshot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsFromAutoSnapshot"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "progress",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Progress"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "size_in_gb",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SizeInGb"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
