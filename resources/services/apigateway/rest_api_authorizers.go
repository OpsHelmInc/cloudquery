// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApiAuthorizers() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_api_authorizers",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html`,
		Resolver:    fetchApigatewayRestApiAuthorizers,
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
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIAuthorizerArn,
			},
			{
				Name:     "auth_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthType"),
			},
			{
				Name:     "authorizer_credentials",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizerCredentials"),
			},
			{
				Name:     "authorizer_result_ttl_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AuthorizerResultTtlInSeconds"),
			},
			{
				Name:     "authorizer_uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizerUri"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "identity_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentitySource"),
			},
			{
				Name:     "identity_validation_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityValidationExpression"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "provider_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ProviderARNs"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
