// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SecurityGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_security_groups",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecurityGroup.html`,
		Resolver:    fetchEc2SecurityGroups,
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
				Resolver: resolveSecurityGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::EC2::SecurityGroup"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupId"),
			},
			{
				Name:     "group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupName"),
			},
			{
				Name:     "ip_permissions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpPermissions"),
			},
			{
				Name:     "ip_permissions_egress",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IpPermissionsEgress"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "security_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityGroupArn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
