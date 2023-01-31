// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayPeeringAttachments() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_transit_gateway_peering_attachments",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html`,
		Resolver:    fetchEc2TransitGatewayPeeringAttachments,
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
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "accepter_tgw_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccepterTgwInfo"),
			},
			{
				Name:     "accepter_transit_gateway_attachment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccepterTransitGatewayAttachmentId"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options"),
			},
			{
				Name:     "requester_tgw_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequesterTgwInfo"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "transit_gateway_attachment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayAttachmentId"),
			},
		},
	}
}
