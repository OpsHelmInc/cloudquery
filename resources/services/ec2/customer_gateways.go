// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func CustomerGateways() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_customer_gateways",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CustomerGateway.html`,
		Resolver:    fetchEc2CustomerGateways,
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
				Resolver: resolveCustomerGatewayArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::EC2::CustomerGateway"),
			},
			{
				Name:     "bgp_asn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BgpAsn"),
			},
			{
				Name:     "certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
			},
			{
				Name:     "customer_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerGatewayId"),
			},
			{
				Name:     "device_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeviceName"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddress"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
