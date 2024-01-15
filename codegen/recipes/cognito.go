package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CognitoResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "identity_pools",
			Struct:              &cognitoidentity.DescribeIdentityPoolOutput{},
			SkipFields:          []string{"IdentityPoolId"},
			PreResourceResolver: "getIdentityPool",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveIdentityPoolARN()`,
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("IdentityPoolId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     ohResourceTypeColumn,
					Type:     schema.TypeString,
					Resolver: `client.StaticValueResolver("AWS::Cognito::IdentityPool")`,
				},
			},
		},
		{
			SubService:          "user_pools",
			Struct:              &types.UserPoolType{},
			Description:         "https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html",
			SkipFields:          []string{"Id"},
			PreResourceResolver: "getUserPool",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: "client.ResolveAWSRegion",
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("Id")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     ohResourceTypeColumn,
					Type:     schema.TypeString,
					Resolver: `client.StaticValueResolver("AWS::Cognito::UserPool")`,
				},
			},
			Relations: []string{
				"UserPoolIdentityProviders()",
			},
		},
		{
			SubService:          "user_pool_identity_providers",
			Struct:              &types.IdentityProviderType{},
			Description:         "https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_IdentityProviderType.html",
			SkipFields:          []string{},
			PreResourceResolver: "getUserPoolIdentityProvider",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "user_pool_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::Cognito::UserPoolIdentityProvider")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cognito"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("cognito-identity")`
	}
	return resources
}
