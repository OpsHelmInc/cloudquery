// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

//go:generate mockgen -package=mocks -destination=../mocks/cloudtrail.go -source=cloudtrail.go CloudtrailClient
type CloudtrailClient interface {
	DescribeQuery(context.Context, *cloudtrail.DescribeQueryInput, ...func(*cloudtrail.Options)) (*cloudtrail.DescribeQueryOutput, error)
	DescribeTrails(context.Context, *cloudtrail.DescribeTrailsInput, ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error)
	GetChannel(context.Context, *cloudtrail.GetChannelInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetChannelOutput, error)
	GetEventDataStore(context.Context, *cloudtrail.GetEventDataStoreInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetEventDataStoreOutput, error)
	GetEventSelectors(context.Context, *cloudtrail.GetEventSelectorsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error)
	GetImport(context.Context, *cloudtrail.GetImportInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetImportOutput, error)
	GetInsightSelectors(context.Context, *cloudtrail.GetInsightSelectorsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetInsightSelectorsOutput, error)
	GetQueryResults(context.Context, *cloudtrail.GetQueryResultsInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetQueryResultsOutput, error)
	GetResourcePolicy(context.Context, *cloudtrail.GetResourcePolicyInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetResourcePolicyOutput, error)
	GetTrail(context.Context, *cloudtrail.GetTrailInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailOutput, error)
	GetTrailStatus(context.Context, *cloudtrail.GetTrailStatusInput, ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error)
	ListChannels(context.Context, *cloudtrail.ListChannelsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListChannelsOutput, error)
	ListEventDataStores(context.Context, *cloudtrail.ListEventDataStoresInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListEventDataStoresOutput, error)
	ListImportFailures(context.Context, *cloudtrail.ListImportFailuresInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListImportFailuresOutput, error)
	ListImports(context.Context, *cloudtrail.ListImportsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListImportsOutput, error)
	ListPublicKeys(context.Context, *cloudtrail.ListPublicKeysInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListPublicKeysOutput, error)
	ListQueries(context.Context, *cloudtrail.ListQueriesInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListQueriesOutput, error)
	ListTags(context.Context, *cloudtrail.ListTagsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error)
	ListTrails(context.Context, *cloudtrail.ListTrailsInput, ...func(*cloudtrail.Options)) (*cloudtrail.ListTrailsOutput, error)
}
