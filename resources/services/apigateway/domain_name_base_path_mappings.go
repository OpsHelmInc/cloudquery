// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DomainNameBasePathMappings() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_domain_name_base_path_mappings",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html`,
		Resolver:    fetchApigatewayDomainNameBasePathMappings,
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
				Name:     "domain_name_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayDomainNameBasePathMappingArn,
			},
			{
				Name:     "base_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BasePath"),
			},
			{
				Name:     "rest_api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RestApiId"),
			},
			{
				Name:     "stage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Stage"),
			},
		},
	}
}
