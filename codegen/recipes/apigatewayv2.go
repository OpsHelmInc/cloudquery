package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func APIGatewayV2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "apis",
			Struct:      &types.Api{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Api.html",
			SkipFields:  []string{"ApiId"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiArn()`,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApiId")`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Api")`,
					},
				}...),
			Relations: []string{
				"ApiAuthorizers()",
				"ApiDeployments()",
				"ApiIntegrations()",
				"ApiModels()",
				"ApiRoutes()",
				"ApiStages()",
			},
		},
		{
			SubService:  "api_authorizers",
			Struct:      &types.Authorizer{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiAuthorizerArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Authorizer")`,
					},
				}...),
		},
		{
			SubService:  "api_deployments",
			Struct:      &types.Deployment{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Deployment.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiDeploymentArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Deployment")`,
					},
				}...),
		},
		{
			SubService:  "api_integrations",
			Struct:      &types.Integration{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Integration.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiIntegrationArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Integration")`,
					},
				}...),
			Relations: []string{
				"ApiIntegrationResponses()",
			},
		},
		{
			SubService:  "api_integration_responses",
			Struct:      &types.IntegrationResponse{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_IntegrationResponse.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_integration_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "integration_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("integration_id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiIntegrationResponseArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::IntegrationResponse")`,
					},
				}...),
		},
		{
			SubService:  "api_models",
			Struct:      &types.Model{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiModelArn()`,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayv2apiModelModelTemplate`,
					},
				}...),
		},
		{
			SubService:  "api_routes",
			Struct:      &types.Route{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Route.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiRouteArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Route")`,
					},
				}...),
			Relations: []string{
				"ApiRouteResponses()",
			},
		},
		{
			SubService:  "api_route_responses",
			Struct:      &types.RouteResponse{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_route_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "route_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("route_id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiRouteResponseArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::RouteResponse")`,
					},
				}...),
		},
		{
			SubService:  "api_stages",
			Struct:      &types.Stage{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiStageArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::Stage")`,
					},
				}...),
		},
		{
			SubService:  "domain_names",
			Struct:      &types.DomainName{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::DomainName")`,
					},
				}...),
			Relations: []string{
				"DomainNameRestApiMappings()",
			},
		},
		{
			SubService:  "domain_name_rest_api_mappings",
			Struct:      &types.ApiMapping{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_ApiMapping.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "domain_name_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameRestApiMappingArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::ApiMapping")`,
					},
				}...),
		},
		{
			SubService:  "vpc_links",
			Struct:      &types.VpcLink{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcLinkArn()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ApiGatewayV2::VpcLink")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "apigatewayv2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("apigateway")`
	}
	return resources
}
