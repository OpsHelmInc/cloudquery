// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LaunchTemplateVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_launch_template_versions",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateVersion.html`,
		Resolver:    fetchEc2LaunchTemplateVersions,
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
				Name:     "launch_template_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchTemplateId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "version_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VersionNumber"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "default_version",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultVersion"),
			},
			{
				Name:     "launch_template_data",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LaunchTemplateData"),
			},
			{
				Name:     "launch_template_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchTemplateName"),
			},
			{
				Name:     "version_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VersionDescription"),
			},
		},
	}
}
