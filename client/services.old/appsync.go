// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

//go:generate mockgen -package=mocks -destination=../mocks/appsync.go -source=appsync.go AppsyncClient
type AppsyncClient interface {
	GetApiAssociation(context.Context, *appsync.GetApiAssociationInput, ...func(*appsync.Options)) (*appsync.GetApiAssociationOutput, error)
	GetApiCache(context.Context, *appsync.GetApiCacheInput, ...func(*appsync.Options)) (*appsync.GetApiCacheOutput, error)
	GetDataSource(context.Context, *appsync.GetDataSourceInput, ...func(*appsync.Options)) (*appsync.GetDataSourceOutput, error)
	GetDataSourceIntrospection(context.Context, *appsync.GetDataSourceIntrospectionInput, ...func(*appsync.Options)) (*appsync.GetDataSourceIntrospectionOutput, error)
	GetDomainName(context.Context, *appsync.GetDomainNameInput, ...func(*appsync.Options)) (*appsync.GetDomainNameOutput, error)
	GetFunction(context.Context, *appsync.GetFunctionInput, ...func(*appsync.Options)) (*appsync.GetFunctionOutput, error)
	GetGraphqlApi(context.Context, *appsync.GetGraphqlApiInput, ...func(*appsync.Options)) (*appsync.GetGraphqlApiOutput, error)
	GetIntrospectionSchema(context.Context, *appsync.GetIntrospectionSchemaInput, ...func(*appsync.Options)) (*appsync.GetIntrospectionSchemaOutput, error)
	GetResolver(context.Context, *appsync.GetResolverInput, ...func(*appsync.Options)) (*appsync.GetResolverOutput, error)
	GetSchemaCreationStatus(context.Context, *appsync.GetSchemaCreationStatusInput, ...func(*appsync.Options)) (*appsync.GetSchemaCreationStatusOutput, error)
	GetSourceApiAssociation(context.Context, *appsync.GetSourceApiAssociationInput, ...func(*appsync.Options)) (*appsync.GetSourceApiAssociationOutput, error)
	GetType(context.Context, *appsync.GetTypeInput, ...func(*appsync.Options)) (*appsync.GetTypeOutput, error)
	ListApiKeys(context.Context, *appsync.ListApiKeysInput, ...func(*appsync.Options)) (*appsync.ListApiKeysOutput, error)
	ListDataSources(context.Context, *appsync.ListDataSourcesInput, ...func(*appsync.Options)) (*appsync.ListDataSourcesOutput, error)
	ListDomainNames(context.Context, *appsync.ListDomainNamesInput, ...func(*appsync.Options)) (*appsync.ListDomainNamesOutput, error)
	ListFunctions(context.Context, *appsync.ListFunctionsInput, ...func(*appsync.Options)) (*appsync.ListFunctionsOutput, error)
	ListGraphqlApis(context.Context, *appsync.ListGraphqlApisInput, ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error)
	ListResolvers(context.Context, *appsync.ListResolversInput, ...func(*appsync.Options)) (*appsync.ListResolversOutput, error)
	ListResolversByFunction(context.Context, *appsync.ListResolversByFunctionInput, ...func(*appsync.Options)) (*appsync.ListResolversByFunctionOutput, error)
	ListSourceApiAssociations(context.Context, *appsync.ListSourceApiAssociationsInput, ...func(*appsync.Options)) (*appsync.ListSourceApiAssociationsOutput, error)
	ListTagsForResource(context.Context, *appsync.ListTagsForResourceInput, ...func(*appsync.Options)) (*appsync.ListTagsForResourceOutput, error)
	ListTypes(context.Context, *appsync.ListTypesInput, ...func(*appsync.Options)) (*appsync.ListTypesOutput, error)
	ListTypesByAssociation(context.Context, *appsync.ListTypesByAssociationInput, ...func(*appsync.Options)) (*appsync.ListTypesByAssociationOutput, error)
}