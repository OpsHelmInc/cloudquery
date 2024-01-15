// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func UsagePlans() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_usage_plans",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlan.html`,
		Resolver:    fetchApigatewayUsagePlans,
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
				Resolver: resolveApigatewayUsagePlanArn,
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::ApiGateway::UsagePlan"),
			},
			{
				Name:     "api_stages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ApiStages"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
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
				Name:     "product_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductCode"),
			},
			{
				Name:     "quota",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Quota"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "throttle",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Throttle"),
			},
		},

		Relations: []*schema.Table{
			UsagePlanKeys(),
		},
	}
}
