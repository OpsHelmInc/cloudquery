// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

//go:generate mockgen -package=mocks -destination=../mocks/redshift.go -source=redshift.go RedshiftClient
type RedshiftClient interface {
	DescribeAccountAttributes(context.Context, *redshift.DescribeAccountAttributesInput, ...func(*redshift.Options)) (*redshift.DescribeAccountAttributesOutput, error)
	DescribeAuthenticationProfiles(context.Context, *redshift.DescribeAuthenticationProfilesInput, ...func(*redshift.Options)) (*redshift.DescribeAuthenticationProfilesOutput, error)
	DescribeClusterDbRevisions(context.Context, *redshift.DescribeClusterDbRevisionsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterDbRevisionsOutput, error)
	DescribeClusterParameterGroups(context.Context, *redshift.DescribeClusterParameterGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterParameterGroupsOutput, error)
	DescribeClusterParameters(context.Context, *redshift.DescribeClusterParametersInput, ...func(*redshift.Options)) (*redshift.DescribeClusterParametersOutput, error)
	DescribeClusterSecurityGroups(context.Context, *redshift.DescribeClusterSecurityGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSecurityGroupsOutput, error)
	DescribeClusterSnapshots(context.Context, *redshift.DescribeClusterSnapshotsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSnapshotsOutput, error)
	DescribeClusterSubnetGroups(context.Context, *redshift.DescribeClusterSubnetGroupsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error)
	DescribeClusterTracks(context.Context, *redshift.DescribeClusterTracksInput, ...func(*redshift.Options)) (*redshift.DescribeClusterTracksOutput, error)
	DescribeClusterVersions(context.Context, *redshift.DescribeClusterVersionsInput, ...func(*redshift.Options)) (*redshift.DescribeClusterVersionsOutput, error)
	DescribeClusters(context.Context, *redshift.DescribeClustersInput, ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error)
	DescribeCustomDomainAssociations(context.Context, *redshift.DescribeCustomDomainAssociationsInput, ...func(*redshift.Options)) (*redshift.DescribeCustomDomainAssociationsOutput, error)
	DescribeDataShares(context.Context, *redshift.DescribeDataSharesInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesOutput, error)
	DescribeDataSharesForConsumer(context.Context, *redshift.DescribeDataSharesForConsumerInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesForConsumerOutput, error)
	DescribeDataSharesForProducer(context.Context, *redshift.DescribeDataSharesForProducerInput, ...func(*redshift.Options)) (*redshift.DescribeDataSharesForProducerOutput, error)
	DescribeDefaultClusterParameters(context.Context, *redshift.DescribeDefaultClusterParametersInput, ...func(*redshift.Options)) (*redshift.DescribeDefaultClusterParametersOutput, error)
	DescribeEndpointAccess(context.Context, *redshift.DescribeEndpointAccessInput, ...func(*redshift.Options)) (*redshift.DescribeEndpointAccessOutput, error)
	DescribeEndpointAuthorization(context.Context, *redshift.DescribeEndpointAuthorizationInput, ...func(*redshift.Options)) (*redshift.DescribeEndpointAuthorizationOutput, error)
	DescribeEventCategories(context.Context, *redshift.DescribeEventCategoriesInput, ...func(*redshift.Options)) (*redshift.DescribeEventCategoriesOutput, error)
	DescribeEventSubscriptions(context.Context, *redshift.DescribeEventSubscriptionsInput, ...func(*redshift.Options)) (*redshift.DescribeEventSubscriptionsOutput, error)
	DescribeEvents(context.Context, *redshift.DescribeEventsInput, ...func(*redshift.Options)) (*redshift.DescribeEventsOutput, error)
	DescribeHsmClientCertificates(context.Context, *redshift.DescribeHsmClientCertificatesInput, ...func(*redshift.Options)) (*redshift.DescribeHsmClientCertificatesOutput, error)
	DescribeHsmConfigurations(context.Context, *redshift.DescribeHsmConfigurationsInput, ...func(*redshift.Options)) (*redshift.DescribeHsmConfigurationsOutput, error)
	DescribeInboundIntegrations(context.Context, *redshift.DescribeInboundIntegrationsInput, ...func(*redshift.Options)) (*redshift.DescribeInboundIntegrationsOutput, error)
	DescribeLoggingStatus(context.Context, *redshift.DescribeLoggingStatusInput, ...func(*redshift.Options)) (*redshift.DescribeLoggingStatusOutput, error)
	DescribeNodeConfigurationOptions(context.Context, *redshift.DescribeNodeConfigurationOptionsInput, ...func(*redshift.Options)) (*redshift.DescribeNodeConfigurationOptionsOutput, error)
	DescribeOrderableClusterOptions(context.Context, *redshift.DescribeOrderableClusterOptionsInput, ...func(*redshift.Options)) (*redshift.DescribeOrderableClusterOptionsOutput, error)
	DescribePartners(context.Context, *redshift.DescribePartnersInput, ...func(*redshift.Options)) (*redshift.DescribePartnersOutput, error)
	DescribeRedshiftIdcApplications(context.Context, *redshift.DescribeRedshiftIdcApplicationsInput, ...func(*redshift.Options)) (*redshift.DescribeRedshiftIdcApplicationsOutput, error)
	DescribeReservedNodeExchangeStatus(context.Context, *redshift.DescribeReservedNodeExchangeStatusInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodeExchangeStatusOutput, error)
	DescribeReservedNodeOfferings(context.Context, *redshift.DescribeReservedNodeOfferingsInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodeOfferingsOutput, error)
	DescribeReservedNodes(context.Context, *redshift.DescribeReservedNodesInput, ...func(*redshift.Options)) (*redshift.DescribeReservedNodesOutput, error)
	DescribeResize(context.Context, *redshift.DescribeResizeInput, ...func(*redshift.Options)) (*redshift.DescribeResizeOutput, error)
	DescribeScheduledActions(context.Context, *redshift.DescribeScheduledActionsInput, ...func(*redshift.Options)) (*redshift.DescribeScheduledActionsOutput, error)
	DescribeSnapshotCopyGrants(context.Context, *redshift.DescribeSnapshotCopyGrantsInput, ...func(*redshift.Options)) (*redshift.DescribeSnapshotCopyGrantsOutput, error)
	DescribeSnapshotSchedules(context.Context, *redshift.DescribeSnapshotSchedulesInput, ...func(*redshift.Options)) (*redshift.DescribeSnapshotSchedulesOutput, error)
	DescribeStorage(context.Context, *redshift.DescribeStorageInput, ...func(*redshift.Options)) (*redshift.DescribeStorageOutput, error)
	DescribeTableRestoreStatus(context.Context, *redshift.DescribeTableRestoreStatusInput, ...func(*redshift.Options)) (*redshift.DescribeTableRestoreStatusOutput, error)
	DescribeTags(context.Context, *redshift.DescribeTagsInput, ...func(*redshift.Options)) (*redshift.DescribeTagsOutput, error)
	DescribeUsageLimits(context.Context, *redshift.DescribeUsageLimitsInput, ...func(*redshift.Options)) (*redshift.DescribeUsageLimitsOutput, error)
	GetClusterCredentials(context.Context, *redshift.GetClusterCredentialsInput, ...func(*redshift.Options)) (*redshift.GetClusterCredentialsOutput, error)
	GetClusterCredentialsWithIAM(context.Context, *redshift.GetClusterCredentialsWithIAMInput, ...func(*redshift.Options)) (*redshift.GetClusterCredentialsWithIAMOutput, error)
	GetReservedNodeExchangeConfigurationOptions(context.Context, *redshift.GetReservedNodeExchangeConfigurationOptionsInput, ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeConfigurationOptionsOutput, error)
	GetReservedNodeExchangeOfferings(context.Context, *redshift.GetReservedNodeExchangeOfferingsInput, ...func(*redshift.Options)) (*redshift.GetReservedNodeExchangeOfferingsOutput, error)
	GetResourcePolicy(context.Context, *redshift.GetResourcePolicyInput, ...func(*redshift.Options)) (*redshift.GetResourcePolicyOutput, error)
	ListRecommendations(context.Context, *redshift.ListRecommendationsInput, ...func(*redshift.Options)) (*redshift.ListRecommendationsOutput, error)
}
