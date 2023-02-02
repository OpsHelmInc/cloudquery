// Code generated by codegen; DO NOT EDIT.

package eventbridge

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiDestinations() *schema.Table {
	return &schema.Table{
		Name:        "aws_eventbridge_api_destinations",
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html`,
		Resolver:    fetchEventbridgeApiDestinations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("events"),
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
				Resolver: schema.PathResolver("ApiDestinationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "api_destination_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiDestinationState"),
			},
			{
				Name:     "connection_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectionArn"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "http_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HttpMethod"),
			},
			{
				Name:     "invocation_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InvocationEndpoint"),
			},
			{
				Name:     "invocation_rate_limit_per_second",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InvocationRateLimitPerSecond"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
