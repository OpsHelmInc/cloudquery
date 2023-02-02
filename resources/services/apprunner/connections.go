// Code generated by codegen; DO NOT EDIT.

package apprunner

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Connections() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_connections",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_Connection.html`,
		Resolver:    fetchApprunnerConnections,
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
				Resolver: schema.PathResolver("ConnectionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("ConnectionArn"),
			},
			{
				Name:     "connection_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionName"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "provider_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProviderType"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
