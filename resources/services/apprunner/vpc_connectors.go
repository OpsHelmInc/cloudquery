// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func VpcConnectors() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_vpc_connectors",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html`,
		Resolver:    fetchApprunnerVpcConnectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apprunner"),
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
				Resolver: schema.PathResolver("VpcConnectorArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("VpcConnectorArn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "deleted_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletedAt"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "vpc_connector_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcConnectorName"),
			},
			{
				Name:     "vpc_connector_revision",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("VpcConnectorRevision"),
			},
		},
	}
}
