// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_images",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html`,
		Resolver:    fetchEc2Images,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::EC2::AMI"),
			},
			{
				Name:     "architecture",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Architecture"),
			},
			{
				Name:     "block_device_mappings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockDeviceMappings"),
			},
			{
				Name:     "boot_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BootMode"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "deprecation_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeprecationTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "ena_support",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnaSupport"),
			},
			{
				Name:     "hypervisor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Hypervisor"),
			},
			{
				Name:     "image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageId"),
			},
			{
				Name:     "image_location",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageLocation"),
			},
			{
				Name:     "image_owner_alias",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageOwnerAlias"),
			},
			{
				Name:     "image_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageType"),
			},
			{
				Name:     "imds_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImdsSupport"),
			},
			{
				Name:     "kernel_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KernelId"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "platform_details",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformDetails"),
			},
			{
				Name:     "product_codes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProductCodes"),
			},
			{
				Name:     "public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Public"),
			},
			{
				Name:     "ramdisk_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RamdiskId"),
			},
			{
				Name:     "root_device_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootDeviceName"),
			},
			{
				Name:     "root_device_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootDeviceType"),
			},
			{
				Name:     "source_instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceInstanceId"),
			},
			{
				Name:     "sriov_net_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SriovNetSupport"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateReason"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tpm_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TpmSupport"),
			},
			{
				Name:     "usage_operation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UsageOperation"),
			},
			{
				Name:     "virtualization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualizationType"),
			},
		},
	}
}
