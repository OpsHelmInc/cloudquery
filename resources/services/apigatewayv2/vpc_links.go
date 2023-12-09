// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func VpcLinks() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_vpc_links",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html`,
		Resolver:    fetchApigatewayv2VpcLinks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Resolver: resolveVpcLinkArn(),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "security_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroupIds"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "vpc_link_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkId"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_link_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkStatus"),
			},
			{
				Name:     "vpc_link_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkStatusMessage"),
			},
			{
				Name:     "vpc_link_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkVersion"),
			},
		},
	}
}
