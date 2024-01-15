// Code generated by codegen; DO NOT EDIT.

package route53

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func HostedZones() *schema.Table {
	return &schema.Table{
		Name:      "aws_route53_hosted_zones",
		Resolver:  fetchRoute53HostedZones,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::Route53::HostedZone"),
			},
			{
				Name:     "caller_reference",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CallerReference"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Config"),
			},
			{
				Name:     "linked_service",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LinkedService"),
			},
			{
				Name:     "resource_record_set_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ResourceRecordSetCount"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "delegation_set_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DelegationSetId"),
			},
			{
				Name:     "vpcs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VPCs"),
			},
		},

		Relations: []*schema.Table{
			HostedZoneQueryLoggingConfigs(),
			HostedZoneResourceRecordSets(),
			HostedZoneTrafficPolicyInstances(),
		},
	}
}
