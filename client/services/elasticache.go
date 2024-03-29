// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

//go:generate mockgen -package=mocks -destination=../mocks/elasticache.go -source=elasticache.go ElasticacheClient
type ElasticacheClient interface {
	DescribeCacheClusters(context.Context, *elasticache.DescribeCacheClustersInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error)
	DescribeCacheEngineVersions(context.Context, *elasticache.DescribeCacheEngineVersionsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error)
	DescribeCacheParameterGroups(context.Context, *elasticache.DescribeCacheParameterGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error)
	DescribeCacheParameters(context.Context, *elasticache.DescribeCacheParametersInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error)
	DescribeCacheSecurityGroups(context.Context, *elasticache.DescribeCacheSecurityGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheSecurityGroupsOutput, error)
	DescribeCacheSubnetGroups(context.Context, *elasticache.DescribeCacheSubnetGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error)
	DescribeEngineDefaultParameters(context.Context, *elasticache.DescribeEngineDefaultParametersInput, ...func(*elasticache.Options)) (*elasticache.DescribeEngineDefaultParametersOutput, error)
	DescribeEvents(context.Context, *elasticache.DescribeEventsInput, ...func(*elasticache.Options)) (*elasticache.DescribeEventsOutput, error)
	DescribeGlobalReplicationGroups(context.Context, *elasticache.DescribeGlobalReplicationGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error)
	DescribeReplicationGroups(context.Context, *elasticache.DescribeReplicationGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error)
	DescribeReservedCacheNodes(context.Context, *elasticache.DescribeReservedCacheNodesInput, ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error)
	DescribeReservedCacheNodesOfferings(context.Context, *elasticache.DescribeReservedCacheNodesOfferingsInput, ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error)
	DescribeServerlessCacheSnapshots(context.Context, *elasticache.DescribeServerlessCacheSnapshotsInput, ...func(*elasticache.Options)) (*elasticache.DescribeServerlessCacheSnapshotsOutput, error)
	DescribeServerlessCaches(context.Context, *elasticache.DescribeServerlessCachesInput, ...func(*elasticache.Options)) (*elasticache.DescribeServerlessCachesOutput, error)
	DescribeServiceUpdates(context.Context, *elasticache.DescribeServiceUpdatesInput, ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error)
	DescribeSnapshots(context.Context, *elasticache.DescribeSnapshotsInput, ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error)
	DescribeUpdateActions(context.Context, *elasticache.DescribeUpdateActionsInput, ...func(*elasticache.Options)) (*elasticache.DescribeUpdateActionsOutput, error)
	DescribeUserGroups(context.Context, *elasticache.DescribeUserGroupsInput, ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error)
	DescribeUsers(context.Context, *elasticache.DescribeUsersInput, ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error)
	ListAllowedNodeTypeModifications(context.Context, *elasticache.ListAllowedNodeTypeModificationsInput, ...func(*elasticache.Options)) (*elasticache.ListAllowedNodeTypeModificationsOutput, error)
	ListTagsForResource(context.Context, *elasticache.ListTagsForResourceInput, ...func(*elasticache.Options)) (*elasticache.ListTagsForResourceOutput, error)
}
