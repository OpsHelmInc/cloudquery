// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

//go:generate mockgen -package=mocks -destination=../mocks/ec2.go -source=ec2.go Ec2Client
type Ec2Client interface {
	DescribeAccountAttributes(context.Context, *ec2.DescribeAccountAttributesInput, ...func(*ec2.Options)) (*ec2.DescribeAccountAttributesOutput, error)
	DescribeAddressTransfers(context.Context, *ec2.DescribeAddressTransfersInput, ...func(*ec2.Options)) (*ec2.DescribeAddressTransfersOutput, error)
	DescribeAddresses(context.Context, *ec2.DescribeAddressesInput, ...func(*ec2.Options)) (*ec2.DescribeAddressesOutput, error)
	DescribeAddressesAttribute(context.Context, *ec2.DescribeAddressesAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeAddressesAttributeOutput, error)
	DescribeAggregateIdFormat(context.Context, *ec2.DescribeAggregateIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeAggregateIdFormatOutput, error)
	DescribeAvailabilityZones(context.Context, *ec2.DescribeAvailabilityZonesInput, ...func(*ec2.Options)) (*ec2.DescribeAvailabilityZonesOutput, error)
	DescribeAwsNetworkPerformanceMetricSubscriptions(context.Context, *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, ...func(*ec2.Options)) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error)
	DescribeBundleTasks(context.Context, *ec2.DescribeBundleTasksInput, ...func(*ec2.Options)) (*ec2.DescribeBundleTasksOutput, error)
	DescribeByoipCidrs(context.Context, *ec2.DescribeByoipCidrsInput, ...func(*ec2.Options)) (*ec2.DescribeByoipCidrsOutput, error)
	DescribeCapacityBlockExtensionHistory(context.Context, *ec2.DescribeCapacityBlockExtensionHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityBlockExtensionHistoryOutput, error)
	DescribeCapacityBlockExtensionOfferings(context.Context, *ec2.DescribeCapacityBlockExtensionOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityBlockExtensionOfferingsOutput, error)
	DescribeCapacityBlockOfferings(context.Context, *ec2.DescribeCapacityBlockOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityBlockOfferingsOutput, error)
	DescribeCapacityReservationBillingRequests(context.Context, *ec2.DescribeCapacityReservationBillingRequestsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationBillingRequestsOutput, error)
	DescribeCapacityReservationFleets(context.Context, *ec2.DescribeCapacityReservationFleetsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationFleetsOutput, error)
	DescribeCapacityReservations(context.Context, *ec2.DescribeCapacityReservationsInput, ...func(*ec2.Options)) (*ec2.DescribeCapacityReservationsOutput, error)
	DescribeCarrierGateways(context.Context, *ec2.DescribeCarrierGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeCarrierGatewaysOutput, error)
	DescribeClassicLinkInstances(context.Context, *ec2.DescribeClassicLinkInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeClassicLinkInstancesOutput, error)
	DescribeClientVpnAuthorizationRules(context.Context, *ec2.DescribeClientVpnAuthorizationRulesInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error)
	DescribeClientVpnConnections(context.Context, *ec2.DescribeClientVpnConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnConnectionsOutput, error)
	DescribeClientVpnEndpoints(context.Context, *ec2.DescribeClientVpnEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnEndpointsOutput, error)
	DescribeClientVpnRoutes(context.Context, *ec2.DescribeClientVpnRoutesInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnRoutesOutput, error)
	DescribeClientVpnTargetNetworks(context.Context, *ec2.DescribeClientVpnTargetNetworksInput, ...func(*ec2.Options)) (*ec2.DescribeClientVpnTargetNetworksOutput, error)
	DescribeCoipPools(context.Context, *ec2.DescribeCoipPoolsInput, ...func(*ec2.Options)) (*ec2.DescribeCoipPoolsOutput, error)
	DescribeConversionTasks(context.Context, *ec2.DescribeConversionTasksInput, ...func(*ec2.Options)) (*ec2.DescribeConversionTasksOutput, error)
	DescribeCustomerGateways(context.Context, *ec2.DescribeCustomerGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeCustomerGatewaysOutput, error)
	DescribeDeclarativePoliciesReports(context.Context, *ec2.DescribeDeclarativePoliciesReportsInput, ...func(*ec2.Options)) (*ec2.DescribeDeclarativePoliciesReportsOutput, error)
	DescribeDhcpOptions(context.Context, *ec2.DescribeDhcpOptionsInput, ...func(*ec2.Options)) (*ec2.DescribeDhcpOptionsOutput, error)
	DescribeEgressOnlyInternetGateways(context.Context, *ec2.DescribeEgressOnlyInternetGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error)
	DescribeElasticGpus(context.Context, *ec2.DescribeElasticGpusInput, ...func(*ec2.Options)) (*ec2.DescribeElasticGpusOutput, error)
	DescribeExportImageTasks(context.Context, *ec2.DescribeExportImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeExportImageTasksOutput, error)
	DescribeExportTasks(context.Context, *ec2.DescribeExportTasksInput, ...func(*ec2.Options)) (*ec2.DescribeExportTasksOutput, error)
	DescribeFastLaunchImages(context.Context, *ec2.DescribeFastLaunchImagesInput, ...func(*ec2.Options)) (*ec2.DescribeFastLaunchImagesOutput, error)
	DescribeFastSnapshotRestores(context.Context, *ec2.DescribeFastSnapshotRestoresInput, ...func(*ec2.Options)) (*ec2.DescribeFastSnapshotRestoresOutput, error)
	DescribeFleetHistory(context.Context, *ec2.DescribeFleetHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeFleetHistoryOutput, error)
	DescribeFleetInstances(context.Context, *ec2.DescribeFleetInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeFleetInstancesOutput, error)
	DescribeFleets(context.Context, *ec2.DescribeFleetsInput, ...func(*ec2.Options)) (*ec2.DescribeFleetsOutput, error)
	DescribeFlowLogs(context.Context, *ec2.DescribeFlowLogsInput, ...func(*ec2.Options)) (*ec2.DescribeFlowLogsOutput, error)
	DescribeFpgaImageAttribute(context.Context, *ec2.DescribeFpgaImageAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeFpgaImageAttributeOutput, error)
	DescribeFpgaImages(context.Context, *ec2.DescribeFpgaImagesInput, ...func(*ec2.Options)) (*ec2.DescribeFpgaImagesOutput, error)
	DescribeHostReservationOfferings(context.Context, *ec2.DescribeHostReservationOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeHostReservationOfferingsOutput, error)
	DescribeHostReservations(context.Context, *ec2.DescribeHostReservationsInput, ...func(*ec2.Options)) (*ec2.DescribeHostReservationsOutput, error)
	DescribeHosts(context.Context, *ec2.DescribeHostsInput, ...func(*ec2.Options)) (*ec2.DescribeHostsOutput, error)
	DescribeIamInstanceProfileAssociations(context.Context, *ec2.DescribeIamInstanceProfileAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error)
	DescribeIdFormat(context.Context, *ec2.DescribeIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeIdFormatOutput, error)
	DescribeIdentityIdFormat(context.Context, *ec2.DescribeIdentityIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribeIdentityIdFormatOutput, error)
	DescribeImageAttribute(context.Context, *ec2.DescribeImageAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeImageAttributeOutput, error)
	DescribeImages(context.Context, *ec2.DescribeImagesInput, ...func(*ec2.Options)) (*ec2.DescribeImagesOutput, error)
	DescribeImportImageTasks(context.Context, *ec2.DescribeImportImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeImportImageTasksOutput, error)
	DescribeImportSnapshotTasks(context.Context, *ec2.DescribeImportSnapshotTasksInput, ...func(*ec2.Options)) (*ec2.DescribeImportSnapshotTasksOutput, error)
	DescribeInstanceAttribute(context.Context, *ec2.DescribeInstanceAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceAttributeOutput, error)
	DescribeInstanceConnectEndpoints(context.Context, *ec2.DescribeInstanceConnectEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceConnectEndpointsOutput, error)
	DescribeInstanceCreditSpecifications(context.Context, *ec2.DescribeInstanceCreditSpecificationsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceCreditSpecificationsOutput, error)
	DescribeInstanceEventNotificationAttributes(context.Context, *ec2.DescribeInstanceEventNotificationAttributesInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error)
	DescribeInstanceEventWindows(context.Context, *ec2.DescribeInstanceEventWindowsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceEventWindowsOutput, error)
	DescribeInstanceImageMetadata(context.Context, *ec2.DescribeInstanceImageMetadataInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceImageMetadataOutput, error)
	DescribeInstanceStatus(context.Context, *ec2.DescribeInstanceStatusInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceStatusOutput, error)
	DescribeInstanceTopology(context.Context, *ec2.DescribeInstanceTopologyInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTopologyOutput, error)
	DescribeInstanceTypeOfferings(context.Context, *ec2.DescribeInstanceTypeOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypeOfferingsOutput, error)
	DescribeInstanceTypes(context.Context, *ec2.DescribeInstanceTypesInput, ...func(*ec2.Options)) (*ec2.DescribeInstanceTypesOutput, error)
	DescribeInstances(context.Context, *ec2.DescribeInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
	DescribeInternetGateways(context.Context, *ec2.DescribeInternetGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeInternetGatewaysOutput, error)
	DescribeIpamByoasn(context.Context, *ec2.DescribeIpamByoasnInput, ...func(*ec2.Options)) (*ec2.DescribeIpamByoasnOutput, error)
	DescribeIpamExternalResourceVerificationTokens(context.Context, *ec2.DescribeIpamExternalResourceVerificationTokensInput, ...func(*ec2.Options)) (*ec2.DescribeIpamExternalResourceVerificationTokensOutput, error)
	DescribeIpamPools(context.Context, *ec2.DescribeIpamPoolsInput, ...func(*ec2.Options)) (*ec2.DescribeIpamPoolsOutput, error)
	DescribeIpamResourceDiscoveries(context.Context, *ec2.DescribeIpamResourceDiscoveriesInput, ...func(*ec2.Options)) (*ec2.DescribeIpamResourceDiscoveriesOutput, error)
	DescribeIpamResourceDiscoveryAssociations(context.Context, *ec2.DescribeIpamResourceDiscoveryAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeIpamResourceDiscoveryAssociationsOutput, error)
	DescribeIpamScopes(context.Context, *ec2.DescribeIpamScopesInput, ...func(*ec2.Options)) (*ec2.DescribeIpamScopesOutput, error)
	DescribeIpams(context.Context, *ec2.DescribeIpamsInput, ...func(*ec2.Options)) (*ec2.DescribeIpamsOutput, error)
	DescribeIpv6Pools(context.Context, *ec2.DescribeIpv6PoolsInput, ...func(*ec2.Options)) (*ec2.DescribeIpv6PoolsOutput, error)
	DescribeKeyPairs(context.Context, *ec2.DescribeKeyPairsInput, ...func(*ec2.Options)) (*ec2.DescribeKeyPairsOutput, error)
	DescribeLaunchTemplateVersions(context.Context, *ec2.DescribeLaunchTemplateVersionsInput, ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplateVersionsOutput, error)
	DescribeLaunchTemplates(context.Context, *ec2.DescribeLaunchTemplatesInput, ...func(*ec2.Options)) (*ec2.DescribeLaunchTemplatesOutput, error)
	DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(context.Context, *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error)
	DescribeLocalGatewayRouteTableVpcAssociations(context.Context, *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error)
	DescribeLocalGatewayRouteTables(context.Context, *ec2.DescribeLocalGatewayRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayRouteTablesOutput, error)
	DescribeLocalGatewayVirtualInterfaceGroups(context.Context, *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error)
	DescribeLocalGatewayVirtualInterfaces(context.Context, *ec2.DescribeLocalGatewayVirtualInterfacesInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error)
	DescribeLocalGateways(context.Context, *ec2.DescribeLocalGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeLocalGatewaysOutput, error)
	DescribeLockedSnapshots(context.Context, *ec2.DescribeLockedSnapshotsInput, ...func(*ec2.Options)) (*ec2.DescribeLockedSnapshotsOutput, error)
	DescribeMacHosts(context.Context, *ec2.DescribeMacHostsInput, ...func(*ec2.Options)) (*ec2.DescribeMacHostsOutput, error)
	DescribeManagedPrefixLists(context.Context, *ec2.DescribeManagedPrefixListsInput, ...func(*ec2.Options)) (*ec2.DescribeManagedPrefixListsOutput, error)
	DescribeMovingAddresses(context.Context, *ec2.DescribeMovingAddressesInput, ...func(*ec2.Options)) (*ec2.DescribeMovingAddressesOutput, error)
	DescribeNatGateways(context.Context, *ec2.DescribeNatGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeNatGatewaysOutput, error)
	DescribeNetworkAcls(context.Context, *ec2.DescribeNetworkAclsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkAclsOutput, error)
	DescribeNetworkInsightsAccessScopeAnalyses(context.Context, *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error)
	DescribeNetworkInsightsAccessScopes(context.Context, *ec2.DescribeNetworkInsightsAccessScopesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error)
	DescribeNetworkInsightsAnalyses(context.Context, *ec2.DescribeNetworkInsightsAnalysesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsAnalysesOutput, error)
	DescribeNetworkInsightsPaths(context.Context, *ec2.DescribeNetworkInsightsPathsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInsightsPathsOutput, error)
	DescribeNetworkInterfaceAttribute(context.Context, *ec2.DescribeNetworkInterfaceAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfaceAttributeOutput, error)
	DescribeNetworkInterfacePermissions(context.Context, *ec2.DescribeNetworkInterfacePermissionsInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacePermissionsOutput, error)
	DescribeNetworkInterfaces(context.Context, *ec2.DescribeNetworkInterfacesInput, ...func(*ec2.Options)) (*ec2.DescribeNetworkInterfacesOutput, error)
	DescribePlacementGroups(context.Context, *ec2.DescribePlacementGroupsInput, ...func(*ec2.Options)) (*ec2.DescribePlacementGroupsOutput, error)
	DescribePrefixLists(context.Context, *ec2.DescribePrefixListsInput, ...func(*ec2.Options)) (*ec2.DescribePrefixListsOutput, error)
	DescribePrincipalIdFormat(context.Context, *ec2.DescribePrincipalIdFormatInput, ...func(*ec2.Options)) (*ec2.DescribePrincipalIdFormatOutput, error)
	DescribePublicIpv4Pools(context.Context, *ec2.DescribePublicIpv4PoolsInput, ...func(*ec2.Options)) (*ec2.DescribePublicIpv4PoolsOutput, error)
	DescribeRegions(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)
	DescribeReplaceRootVolumeTasks(context.Context, *ec2.DescribeReplaceRootVolumeTasksInput, ...func(*ec2.Options)) (*ec2.DescribeReplaceRootVolumeTasksOutput, error)
	DescribeReservedInstances(context.Context, *ec2.DescribeReservedInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOutput, error)
	DescribeReservedInstancesListings(context.Context, *ec2.DescribeReservedInstancesListingsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesListingsOutput, error)
	DescribeReservedInstancesModifications(context.Context, *ec2.DescribeReservedInstancesModificationsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesModificationsOutput, error)
	DescribeReservedInstancesOfferings(context.Context, *ec2.DescribeReservedInstancesOfferingsInput, ...func(*ec2.Options)) (*ec2.DescribeReservedInstancesOfferingsOutput, error)
	DescribeRouteTables(context.Context, *ec2.DescribeRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeRouteTablesOutput, error)
	DescribeScheduledInstanceAvailability(context.Context, *ec2.DescribeScheduledInstanceAvailabilityInput, ...func(*ec2.Options)) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error)
	DescribeScheduledInstances(context.Context, *ec2.DescribeScheduledInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeScheduledInstancesOutput, error)
	DescribeSecurityGroupReferences(context.Context, *ec2.DescribeSecurityGroupReferencesInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupReferencesOutput, error)
	DescribeSecurityGroupRules(context.Context, *ec2.DescribeSecurityGroupRulesInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupRulesOutput, error)
	DescribeSecurityGroupVpcAssociations(context.Context, *ec2.DescribeSecurityGroupVpcAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupVpcAssociationsOutput, error)
	DescribeSecurityGroups(context.Context, *ec2.DescribeSecurityGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeSecurityGroupsOutput, error)
	DescribeSnapshotAttribute(context.Context, *ec2.DescribeSnapshotAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotAttributeOutput, error)
	DescribeSnapshotTierStatus(context.Context, *ec2.DescribeSnapshotTierStatusInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotTierStatusOutput, error)
	DescribeSnapshots(context.Context, *ec2.DescribeSnapshotsInput, ...func(*ec2.Options)) (*ec2.DescribeSnapshotsOutput, error)
	DescribeSpotDatafeedSubscription(context.Context, *ec2.DescribeSpotDatafeedSubscriptionInput, ...func(*ec2.Options)) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error)
	DescribeSpotFleetInstances(context.Context, *ec2.DescribeSpotFleetInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetInstancesOutput, error)
	DescribeSpotFleetRequestHistory(context.Context, *ec2.DescribeSpotFleetRequestHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestHistoryOutput, error)
	DescribeSpotFleetRequests(context.Context, *ec2.DescribeSpotFleetRequestsInput, ...func(*ec2.Options)) (*ec2.DescribeSpotFleetRequestsOutput, error)
	DescribeSpotInstanceRequests(context.Context, *ec2.DescribeSpotInstanceRequestsInput, ...func(*ec2.Options)) (*ec2.DescribeSpotInstanceRequestsOutput, error)
	DescribeSpotPriceHistory(context.Context, *ec2.DescribeSpotPriceHistoryInput, ...func(*ec2.Options)) (*ec2.DescribeSpotPriceHistoryOutput, error)
	DescribeStaleSecurityGroups(context.Context, *ec2.DescribeStaleSecurityGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeStaleSecurityGroupsOutput, error)
	DescribeStoreImageTasks(context.Context, *ec2.DescribeStoreImageTasksInput, ...func(*ec2.Options)) (*ec2.DescribeStoreImageTasksOutput, error)
	DescribeSubnets(context.Context, *ec2.DescribeSubnetsInput, ...func(*ec2.Options)) (*ec2.DescribeSubnetsOutput, error)
	DescribeTags(context.Context, *ec2.DescribeTagsInput, ...func(*ec2.Options)) (*ec2.DescribeTagsOutput, error)
	DescribeTrafficMirrorFilterRules(context.Context, *ec2.DescribeTrafficMirrorFilterRulesInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFilterRulesOutput, error)
	DescribeTrafficMirrorFilters(context.Context, *ec2.DescribeTrafficMirrorFiltersInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorFiltersOutput, error)
	DescribeTrafficMirrorSessions(context.Context, *ec2.DescribeTrafficMirrorSessionsInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorSessionsOutput, error)
	DescribeTrafficMirrorTargets(context.Context, *ec2.DescribeTrafficMirrorTargetsInput, ...func(*ec2.Options)) (*ec2.DescribeTrafficMirrorTargetsOutput, error)
	DescribeTransitGatewayAttachments(context.Context, *ec2.DescribeTransitGatewayAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayAttachmentsOutput, error)
	DescribeTransitGatewayConnectPeers(context.Context, *ec2.DescribeTransitGatewayConnectPeersInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectPeersOutput, error)
	DescribeTransitGatewayConnects(context.Context, *ec2.DescribeTransitGatewayConnectsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayConnectsOutput, error)
	DescribeTransitGatewayMulticastDomains(context.Context, *ec2.DescribeTransitGatewayMulticastDomainsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error)
	DescribeTransitGatewayPeeringAttachments(context.Context, *ec2.DescribeTransitGatewayPeeringAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error)
	DescribeTransitGatewayPolicyTables(context.Context, *ec2.DescribeTransitGatewayPolicyTablesInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error)
	DescribeTransitGatewayRouteTableAnnouncements(context.Context, *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error)
	DescribeTransitGatewayRouteTables(context.Context, *ec2.DescribeTransitGatewayRouteTablesInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayRouteTablesOutput, error)
	DescribeTransitGatewayVpcAttachments(context.Context, *ec2.DescribeTransitGatewayVpcAttachmentsInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error)
	DescribeTransitGateways(context.Context, *ec2.DescribeTransitGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeTransitGatewaysOutput, error)
	DescribeTrunkInterfaceAssociations(context.Context, *ec2.DescribeTrunkInterfaceAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error)
	DescribeVerifiedAccessEndpoints(context.Context, *ec2.DescribeVerifiedAccessEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessEndpointsOutput, error)
	DescribeVerifiedAccessGroups(context.Context, *ec2.DescribeVerifiedAccessGroupsInput, ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessGroupsOutput, error)
	DescribeVerifiedAccessInstanceLoggingConfigurations(context.Context, *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error)
	DescribeVerifiedAccessInstances(context.Context, *ec2.DescribeVerifiedAccessInstancesInput, ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessInstancesOutput, error)
	DescribeVerifiedAccessTrustProviders(context.Context, *ec2.DescribeVerifiedAccessTrustProvidersInput, ...func(*ec2.Options)) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error)
	DescribeVolumeAttribute(context.Context, *ec2.DescribeVolumeAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeVolumeAttributeOutput, error)
	DescribeVolumeStatus(context.Context, *ec2.DescribeVolumeStatusInput, ...func(*ec2.Options)) (*ec2.DescribeVolumeStatusOutput, error)
	DescribeVolumes(context.Context, *ec2.DescribeVolumesInput, ...func(*ec2.Options)) (*ec2.DescribeVolumesOutput, error)
	DescribeVolumesModifications(context.Context, *ec2.DescribeVolumesModificationsInput, ...func(*ec2.Options)) (*ec2.DescribeVolumesModificationsOutput, error)
	DescribeVpcAttribute(context.Context, *ec2.DescribeVpcAttributeInput, ...func(*ec2.Options)) (*ec2.DescribeVpcAttributeOutput, error)
	DescribeVpcBlockPublicAccessExclusions(context.Context, *ec2.DescribeVpcBlockPublicAccessExclusionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcBlockPublicAccessExclusionsOutput, error)
	DescribeVpcBlockPublicAccessOptions(context.Context, *ec2.DescribeVpcBlockPublicAccessOptionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcBlockPublicAccessOptionsOutput, error)
	DescribeVpcClassicLink(context.Context, *ec2.DescribeVpcClassicLinkInput, ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkOutput, error)
	DescribeVpcClassicLinkDnsSupport(context.Context, *ec2.DescribeVpcClassicLinkDnsSupportInput, ...func(*ec2.Options)) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error)
	DescribeVpcEndpointAssociations(context.Context, *ec2.DescribeVpcEndpointAssociationsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointAssociationsOutput, error)
	DescribeVpcEndpointConnectionNotifications(context.Context, *ec2.DescribeVpcEndpointConnectionNotificationsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error)
	DescribeVpcEndpointConnections(context.Context, *ec2.DescribeVpcEndpointConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointConnectionsOutput, error)
	DescribeVpcEndpointServiceConfigurations(context.Context, *ec2.DescribeVpcEndpointServiceConfigurationsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error)
	DescribeVpcEndpointServicePermissions(context.Context, *ec2.DescribeVpcEndpointServicePermissionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error)
	DescribeVpcEndpointServices(context.Context, *ec2.DescribeVpcEndpointServicesInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointServicesOutput, error)
	DescribeVpcEndpoints(context.Context, *ec2.DescribeVpcEndpointsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcEndpointsOutput, error)
	DescribeVpcPeeringConnections(context.Context, *ec2.DescribeVpcPeeringConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcPeeringConnectionsOutput, error)
	DescribeVpcs(context.Context, *ec2.DescribeVpcsInput, ...func(*ec2.Options)) (*ec2.DescribeVpcsOutput, error)
	DescribeVpnConnections(context.Context, *ec2.DescribeVpnConnectionsInput, ...func(*ec2.Options)) (*ec2.DescribeVpnConnectionsOutput, error)
	DescribeVpnGateways(context.Context, *ec2.DescribeVpnGatewaysInput, ...func(*ec2.Options)) (*ec2.DescribeVpnGatewaysOutput, error)
	GetAllowedImagesSettings(context.Context, *ec2.GetAllowedImagesSettingsInput, ...func(*ec2.Options)) (*ec2.GetAllowedImagesSettingsOutput, error)
	GetAssociatedEnclaveCertificateIamRoles(context.Context, *ec2.GetAssociatedEnclaveCertificateIamRolesInput, ...func(*ec2.Options)) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error)
	GetAssociatedIpv6PoolCidrs(context.Context, *ec2.GetAssociatedIpv6PoolCidrsInput, ...func(*ec2.Options)) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error)
	GetAwsNetworkPerformanceData(context.Context, *ec2.GetAwsNetworkPerformanceDataInput, ...func(*ec2.Options)) (*ec2.GetAwsNetworkPerformanceDataOutput, error)
	GetCapacityReservationUsage(context.Context, *ec2.GetCapacityReservationUsageInput, ...func(*ec2.Options)) (*ec2.GetCapacityReservationUsageOutput, error)
	GetCoipPoolUsage(context.Context, *ec2.GetCoipPoolUsageInput, ...func(*ec2.Options)) (*ec2.GetCoipPoolUsageOutput, error)
	GetConsoleOutput(context.Context, *ec2.GetConsoleOutputInput, ...func(*ec2.Options)) (*ec2.GetConsoleOutputOutput, error)
	GetConsoleScreenshot(context.Context, *ec2.GetConsoleScreenshotInput, ...func(*ec2.Options)) (*ec2.GetConsoleScreenshotOutput, error)
	GetDeclarativePoliciesReportSummary(context.Context, *ec2.GetDeclarativePoliciesReportSummaryInput, ...func(*ec2.Options)) (*ec2.GetDeclarativePoliciesReportSummaryOutput, error)
	GetDefaultCreditSpecification(context.Context, *ec2.GetDefaultCreditSpecificationInput, ...func(*ec2.Options)) (*ec2.GetDefaultCreditSpecificationOutput, error)
	GetEbsDefaultKmsKeyId(context.Context, *ec2.GetEbsDefaultKmsKeyIdInput, ...func(*ec2.Options)) (*ec2.GetEbsDefaultKmsKeyIdOutput, error)
	GetEbsEncryptionByDefault(context.Context, *ec2.GetEbsEncryptionByDefaultInput, ...func(*ec2.Options)) (*ec2.GetEbsEncryptionByDefaultOutput, error)
	GetFlowLogsIntegrationTemplate(context.Context, *ec2.GetFlowLogsIntegrationTemplateInput, ...func(*ec2.Options)) (*ec2.GetFlowLogsIntegrationTemplateOutput, error)
	GetGroupsForCapacityReservation(context.Context, *ec2.GetGroupsForCapacityReservationInput, ...func(*ec2.Options)) (*ec2.GetGroupsForCapacityReservationOutput, error)
	GetHostReservationPurchasePreview(context.Context, *ec2.GetHostReservationPurchasePreviewInput, ...func(*ec2.Options)) (*ec2.GetHostReservationPurchasePreviewOutput, error)
	GetImageBlockPublicAccessState(context.Context, *ec2.GetImageBlockPublicAccessStateInput, ...func(*ec2.Options)) (*ec2.GetImageBlockPublicAccessStateOutput, error)
	GetInstanceMetadataDefaults(context.Context, *ec2.GetInstanceMetadataDefaultsInput, ...func(*ec2.Options)) (*ec2.GetInstanceMetadataDefaultsOutput, error)
	GetInstanceTpmEkPub(context.Context, *ec2.GetInstanceTpmEkPubInput, ...func(*ec2.Options)) (*ec2.GetInstanceTpmEkPubOutput, error)
	GetInstanceTypesFromInstanceRequirements(context.Context, *ec2.GetInstanceTypesFromInstanceRequirementsInput, ...func(*ec2.Options)) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error)
	GetInstanceUefiData(context.Context, *ec2.GetInstanceUefiDataInput, ...func(*ec2.Options)) (*ec2.GetInstanceUefiDataOutput, error)
	GetIpamAddressHistory(context.Context, *ec2.GetIpamAddressHistoryInput, ...func(*ec2.Options)) (*ec2.GetIpamAddressHistoryOutput, error)
	GetIpamDiscoveredAccounts(context.Context, *ec2.GetIpamDiscoveredAccountsInput, ...func(*ec2.Options)) (*ec2.GetIpamDiscoveredAccountsOutput, error)
	GetIpamDiscoveredPublicAddresses(context.Context, *ec2.GetIpamDiscoveredPublicAddressesInput, ...func(*ec2.Options)) (*ec2.GetIpamDiscoveredPublicAddressesOutput, error)
	GetIpamDiscoveredResourceCidrs(context.Context, *ec2.GetIpamDiscoveredResourceCidrsInput, ...func(*ec2.Options)) (*ec2.GetIpamDiscoveredResourceCidrsOutput, error)
	GetIpamPoolAllocations(context.Context, *ec2.GetIpamPoolAllocationsInput, ...func(*ec2.Options)) (*ec2.GetIpamPoolAllocationsOutput, error)
	GetIpamPoolCidrs(context.Context, *ec2.GetIpamPoolCidrsInput, ...func(*ec2.Options)) (*ec2.GetIpamPoolCidrsOutput, error)
	GetIpamResourceCidrs(context.Context, *ec2.GetIpamResourceCidrsInput, ...func(*ec2.Options)) (*ec2.GetIpamResourceCidrsOutput, error)
	GetLaunchTemplateData(context.Context, *ec2.GetLaunchTemplateDataInput, ...func(*ec2.Options)) (*ec2.GetLaunchTemplateDataOutput, error)
	GetManagedPrefixListAssociations(context.Context, *ec2.GetManagedPrefixListAssociationsInput, ...func(*ec2.Options)) (*ec2.GetManagedPrefixListAssociationsOutput, error)
	GetManagedPrefixListEntries(context.Context, *ec2.GetManagedPrefixListEntriesInput, ...func(*ec2.Options)) (*ec2.GetManagedPrefixListEntriesOutput, error)
	GetNetworkInsightsAccessScopeAnalysisFindings(context.Context, *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error)
	GetNetworkInsightsAccessScopeContent(context.Context, *ec2.GetNetworkInsightsAccessScopeContentInput, ...func(*ec2.Options)) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error)
	GetPasswordData(context.Context, *ec2.GetPasswordDataInput, ...func(*ec2.Options)) (*ec2.GetPasswordDataOutput, error)
	GetReservedInstancesExchangeQuote(context.Context, *ec2.GetReservedInstancesExchangeQuoteInput, ...func(*ec2.Options)) (*ec2.GetReservedInstancesExchangeQuoteOutput, error)
	GetSecurityGroupsForVpc(context.Context, *ec2.GetSecurityGroupsForVpcInput, ...func(*ec2.Options)) (*ec2.GetSecurityGroupsForVpcOutput, error)
	GetSerialConsoleAccessStatus(context.Context, *ec2.GetSerialConsoleAccessStatusInput, ...func(*ec2.Options)) (*ec2.GetSerialConsoleAccessStatusOutput, error)
	GetSnapshotBlockPublicAccessState(context.Context, *ec2.GetSnapshotBlockPublicAccessStateInput, ...func(*ec2.Options)) (*ec2.GetSnapshotBlockPublicAccessStateOutput, error)
	GetSpotPlacementScores(context.Context, *ec2.GetSpotPlacementScoresInput, ...func(*ec2.Options)) (*ec2.GetSpotPlacementScoresOutput, error)
	GetSubnetCidrReservations(context.Context, *ec2.GetSubnetCidrReservationsInput, ...func(*ec2.Options)) (*ec2.GetSubnetCidrReservationsOutput, error)
	GetTransitGatewayAttachmentPropagations(context.Context, *ec2.GetTransitGatewayAttachmentPropagationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error)
	GetTransitGatewayMulticastDomainAssociations(context.Context, *ec2.GetTransitGatewayMulticastDomainAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error)
	GetTransitGatewayPolicyTableAssociations(context.Context, *ec2.GetTransitGatewayPolicyTableAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error)
	GetTransitGatewayPolicyTableEntries(context.Context, *ec2.GetTransitGatewayPolicyTableEntriesInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error)
	GetTransitGatewayPrefixListReferences(context.Context, *ec2.GetTransitGatewayPrefixListReferencesInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error)
	GetTransitGatewayRouteTableAssociations(context.Context, *ec2.GetTransitGatewayRouteTableAssociationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error)
	GetTransitGatewayRouteTablePropagations(context.Context, *ec2.GetTransitGatewayRouteTablePropagationsInput, ...func(*ec2.Options)) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error)
	GetVerifiedAccessEndpointPolicy(context.Context, *ec2.GetVerifiedAccessEndpointPolicyInput, ...func(*ec2.Options)) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error)
	GetVerifiedAccessEndpointTargets(context.Context, *ec2.GetVerifiedAccessEndpointTargetsInput, ...func(*ec2.Options)) (*ec2.GetVerifiedAccessEndpointTargetsOutput, error)
	GetVerifiedAccessGroupPolicy(context.Context, *ec2.GetVerifiedAccessGroupPolicyInput, ...func(*ec2.Options)) (*ec2.GetVerifiedAccessGroupPolicyOutput, error)
	GetVpnConnectionDeviceSampleConfiguration(context.Context, *ec2.GetVpnConnectionDeviceSampleConfigurationInput, ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error)
	GetVpnConnectionDeviceTypes(context.Context, *ec2.GetVpnConnectionDeviceTypesInput, ...func(*ec2.Options)) (*ec2.GetVpnConnectionDeviceTypesOutput, error)
	GetVpnTunnelReplacementStatus(context.Context, *ec2.GetVpnTunnelReplacementStatusInput, ...func(*ec2.Options)) (*ec2.GetVpnTunnelReplacementStatusOutput, error)
	ListImagesInRecycleBin(context.Context, *ec2.ListImagesInRecycleBinInput, ...func(*ec2.Options)) (*ec2.ListImagesInRecycleBinOutput, error)
	ListSnapshotsInRecycleBin(context.Context, *ec2.ListSnapshotsInRecycleBinInput, ...func(*ec2.Options)) (*ec2.ListSnapshotsInRecycleBinOutput, error)
	SearchLocalGatewayRoutes(context.Context, *ec2.SearchLocalGatewayRoutesInput, ...func(*ec2.Options)) (*ec2.SearchLocalGatewayRoutesOutput, error)
	SearchTransitGatewayMulticastGroups(context.Context, *ec2.SearchTransitGatewayMulticastGroupsInput, ...func(*ec2.Options)) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error)
	SearchTransitGatewayRoutes(context.Context, *ec2.SearchTransitGatewayRoutesInput, ...func(*ec2.Options)) (*ec2.SearchTransitGatewayRoutesOutput, error)
}
