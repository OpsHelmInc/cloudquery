// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
)

//go:generate mockgen -package=mocks -destination=../mocks/quicksight.go -source=quicksight.go QuicksightClient
type QuicksightClient interface {
	DescribeAccountCustomization(context.Context, *quicksight.DescribeAccountCustomizationInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountCustomizationOutput, error)
	DescribeAccountSettings(context.Context, *quicksight.DescribeAccountSettingsInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountSettingsOutput, error)
	DescribeAccountSubscription(context.Context, *quicksight.DescribeAccountSubscriptionInput, ...func(*quicksight.Options)) (*quicksight.DescribeAccountSubscriptionOutput, error)
	DescribeAnalysis(context.Context, *quicksight.DescribeAnalysisInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisOutput, error)
	DescribeAnalysisDefinition(context.Context, *quicksight.DescribeAnalysisDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisDefinitionOutput, error)
	DescribeAnalysisPermissions(context.Context, *quicksight.DescribeAnalysisPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeAnalysisPermissionsOutput, error)
	DescribeAssetBundleExportJob(context.Context, *quicksight.DescribeAssetBundleExportJobInput, ...func(*quicksight.Options)) (*quicksight.DescribeAssetBundleExportJobOutput, error)
	DescribeAssetBundleImportJob(context.Context, *quicksight.DescribeAssetBundleImportJobInput, ...func(*quicksight.Options)) (*quicksight.DescribeAssetBundleImportJobOutput, error)
	DescribeDashboard(context.Context, *quicksight.DescribeDashboardInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardOutput, error)
	DescribeDashboardDefinition(context.Context, *quicksight.DescribeDashboardDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardDefinitionOutput, error)
	DescribeDashboardPermissions(context.Context, *quicksight.DescribeDashboardPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardPermissionsOutput, error)
	DescribeDashboardSnapshotJob(context.Context, *quicksight.DescribeDashboardSnapshotJobInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardSnapshotJobOutput, error)
	DescribeDashboardSnapshotJobResult(context.Context, *quicksight.DescribeDashboardSnapshotJobResultInput, ...func(*quicksight.Options)) (*quicksight.DescribeDashboardSnapshotJobResultOutput, error)
	DescribeDataSet(context.Context, *quicksight.DescribeDataSetInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSetOutput, error)
	DescribeDataSetPermissions(context.Context, *quicksight.DescribeDataSetPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSetPermissionsOutput, error)
	DescribeDataSetRefreshProperties(context.Context, *quicksight.DescribeDataSetRefreshPropertiesInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSetRefreshPropertiesOutput, error)
	DescribeDataSource(context.Context, *quicksight.DescribeDataSourceInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSourceOutput, error)
	DescribeDataSourcePermissions(context.Context, *quicksight.DescribeDataSourcePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeDataSourcePermissionsOutput, error)
	DescribeFolder(context.Context, *quicksight.DescribeFolderInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderOutput, error)
	DescribeFolderPermissions(context.Context, *quicksight.DescribeFolderPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderPermissionsOutput, error)
	DescribeFolderResolvedPermissions(context.Context, *quicksight.DescribeFolderResolvedPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeFolderResolvedPermissionsOutput, error)
	DescribeGroup(context.Context, *quicksight.DescribeGroupInput, ...func(*quicksight.Options)) (*quicksight.DescribeGroupOutput, error)
	DescribeGroupMembership(context.Context, *quicksight.DescribeGroupMembershipInput, ...func(*quicksight.Options)) (*quicksight.DescribeGroupMembershipOutput, error)
	DescribeIAMPolicyAssignment(context.Context, *quicksight.DescribeIAMPolicyAssignmentInput, ...func(*quicksight.Options)) (*quicksight.DescribeIAMPolicyAssignmentOutput, error)
	DescribeIngestion(context.Context, *quicksight.DescribeIngestionInput, ...func(*quicksight.Options)) (*quicksight.DescribeIngestionOutput, error)
	DescribeIpRestriction(context.Context, *quicksight.DescribeIpRestrictionInput, ...func(*quicksight.Options)) (*quicksight.DescribeIpRestrictionOutput, error)
	DescribeNamespace(context.Context, *quicksight.DescribeNamespaceInput, ...func(*quicksight.Options)) (*quicksight.DescribeNamespaceOutput, error)
	DescribeRefreshSchedule(context.Context, *quicksight.DescribeRefreshScheduleInput, ...func(*quicksight.Options)) (*quicksight.DescribeRefreshScheduleOutput, error)
	DescribeRoleCustomPermission(context.Context, *quicksight.DescribeRoleCustomPermissionInput, ...func(*quicksight.Options)) (*quicksight.DescribeRoleCustomPermissionOutput, error)
	DescribeTemplate(context.Context, *quicksight.DescribeTemplateInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateOutput, error)
	DescribeTemplateAlias(context.Context, *quicksight.DescribeTemplateAliasInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateAliasOutput, error)
	DescribeTemplateDefinition(context.Context, *quicksight.DescribeTemplateDefinitionInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplateDefinitionOutput, error)
	DescribeTemplatePermissions(context.Context, *quicksight.DescribeTemplatePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeTemplatePermissionsOutput, error)
	DescribeTheme(context.Context, *quicksight.DescribeThemeInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemeOutput, error)
	DescribeThemeAlias(context.Context, *quicksight.DescribeThemeAliasInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemeAliasOutput, error)
	DescribeThemePermissions(context.Context, *quicksight.DescribeThemePermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeThemePermissionsOutput, error)
	DescribeTopic(context.Context, *quicksight.DescribeTopicInput, ...func(*quicksight.Options)) (*quicksight.DescribeTopicOutput, error)
	DescribeTopicPermissions(context.Context, *quicksight.DescribeTopicPermissionsInput, ...func(*quicksight.Options)) (*quicksight.DescribeTopicPermissionsOutput, error)
	DescribeTopicRefresh(context.Context, *quicksight.DescribeTopicRefreshInput, ...func(*quicksight.Options)) (*quicksight.DescribeTopicRefreshOutput, error)
	DescribeTopicRefreshSchedule(context.Context, *quicksight.DescribeTopicRefreshScheduleInput, ...func(*quicksight.Options)) (*quicksight.DescribeTopicRefreshScheduleOutput, error)
	DescribeUser(context.Context, *quicksight.DescribeUserInput, ...func(*quicksight.Options)) (*quicksight.DescribeUserOutput, error)
	DescribeVPCConnection(context.Context, *quicksight.DescribeVPCConnectionInput, ...func(*quicksight.Options)) (*quicksight.DescribeVPCConnectionOutput, error)
	GetDashboardEmbedUrl(context.Context, *quicksight.GetDashboardEmbedUrlInput, ...func(*quicksight.Options)) (*quicksight.GetDashboardEmbedUrlOutput, error)
	GetSessionEmbedUrl(context.Context, *quicksight.GetSessionEmbedUrlInput, ...func(*quicksight.Options)) (*quicksight.GetSessionEmbedUrlOutput, error)
	ListAnalyses(context.Context, *quicksight.ListAnalysesInput, ...func(*quicksight.Options)) (*quicksight.ListAnalysesOutput, error)
	ListAssetBundleExportJobs(context.Context, *quicksight.ListAssetBundleExportJobsInput, ...func(*quicksight.Options)) (*quicksight.ListAssetBundleExportJobsOutput, error)
	ListAssetBundleImportJobs(context.Context, *quicksight.ListAssetBundleImportJobsInput, ...func(*quicksight.Options)) (*quicksight.ListAssetBundleImportJobsOutput, error)
	ListDashboardVersions(context.Context, *quicksight.ListDashboardVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListDashboardVersionsOutput, error)
	ListDashboards(context.Context, *quicksight.ListDashboardsInput, ...func(*quicksight.Options)) (*quicksight.ListDashboardsOutput, error)
	ListDataSets(context.Context, *quicksight.ListDataSetsInput, ...func(*quicksight.Options)) (*quicksight.ListDataSetsOutput, error)
	ListDataSources(context.Context, *quicksight.ListDataSourcesInput, ...func(*quicksight.Options)) (*quicksight.ListDataSourcesOutput, error)
	ListFolderMembers(context.Context, *quicksight.ListFolderMembersInput, ...func(*quicksight.Options)) (*quicksight.ListFolderMembersOutput, error)
	ListFolders(context.Context, *quicksight.ListFoldersInput, ...func(*quicksight.Options)) (*quicksight.ListFoldersOutput, error)
	ListGroupMemberships(context.Context, *quicksight.ListGroupMembershipsInput, ...func(*quicksight.Options)) (*quicksight.ListGroupMembershipsOutput, error)
	ListGroups(context.Context, *quicksight.ListGroupsInput, ...func(*quicksight.Options)) (*quicksight.ListGroupsOutput, error)
	ListIAMPolicyAssignments(context.Context, *quicksight.ListIAMPolicyAssignmentsInput, ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsOutput, error)
	ListIAMPolicyAssignmentsForUser(context.Context, *quicksight.ListIAMPolicyAssignmentsForUserInput, ...func(*quicksight.Options)) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error)
	ListIdentityPropagationConfigs(context.Context, *quicksight.ListIdentityPropagationConfigsInput, ...func(*quicksight.Options)) (*quicksight.ListIdentityPropagationConfigsOutput, error)
	ListIngestions(context.Context, *quicksight.ListIngestionsInput, ...func(*quicksight.Options)) (*quicksight.ListIngestionsOutput, error)
	ListNamespaces(context.Context, *quicksight.ListNamespacesInput, ...func(*quicksight.Options)) (*quicksight.ListNamespacesOutput, error)
	ListRefreshSchedules(context.Context, *quicksight.ListRefreshSchedulesInput, ...func(*quicksight.Options)) (*quicksight.ListRefreshSchedulesOutput, error)
	ListRoleMemberships(context.Context, *quicksight.ListRoleMembershipsInput, ...func(*quicksight.Options)) (*quicksight.ListRoleMembershipsOutput, error)
	ListTagsForResource(context.Context, *quicksight.ListTagsForResourceInput, ...func(*quicksight.Options)) (*quicksight.ListTagsForResourceOutput, error)
	ListTemplateAliases(context.Context, *quicksight.ListTemplateAliasesInput, ...func(*quicksight.Options)) (*quicksight.ListTemplateAliasesOutput, error)
	ListTemplateVersions(context.Context, *quicksight.ListTemplateVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListTemplateVersionsOutput, error)
	ListTemplates(context.Context, *quicksight.ListTemplatesInput, ...func(*quicksight.Options)) (*quicksight.ListTemplatesOutput, error)
	ListThemeAliases(context.Context, *quicksight.ListThemeAliasesInput, ...func(*quicksight.Options)) (*quicksight.ListThemeAliasesOutput, error)
	ListThemeVersions(context.Context, *quicksight.ListThemeVersionsInput, ...func(*quicksight.Options)) (*quicksight.ListThemeVersionsOutput, error)
	ListThemes(context.Context, *quicksight.ListThemesInput, ...func(*quicksight.Options)) (*quicksight.ListThemesOutput, error)
	ListTopicRefreshSchedules(context.Context, *quicksight.ListTopicRefreshSchedulesInput, ...func(*quicksight.Options)) (*quicksight.ListTopicRefreshSchedulesOutput, error)
	ListTopics(context.Context, *quicksight.ListTopicsInput, ...func(*quicksight.Options)) (*quicksight.ListTopicsOutput, error)
	ListUserGroups(context.Context, *quicksight.ListUserGroupsInput, ...func(*quicksight.Options)) (*quicksight.ListUserGroupsOutput, error)
	ListUsers(context.Context, *quicksight.ListUsersInput, ...func(*quicksight.Options)) (*quicksight.ListUsersOutput, error)
	ListVPCConnections(context.Context, *quicksight.ListVPCConnectionsInput, ...func(*quicksight.Options)) (*quicksight.ListVPCConnectionsOutput, error)
	SearchAnalyses(context.Context, *quicksight.SearchAnalysesInput, ...func(*quicksight.Options)) (*quicksight.SearchAnalysesOutput, error)
	SearchDashboards(context.Context, *quicksight.SearchDashboardsInput, ...func(*quicksight.Options)) (*quicksight.SearchDashboardsOutput, error)
	SearchDataSets(context.Context, *quicksight.SearchDataSetsInput, ...func(*quicksight.Options)) (*quicksight.SearchDataSetsOutput, error)
	SearchDataSources(context.Context, *quicksight.SearchDataSourcesInput, ...func(*quicksight.Options)) (*quicksight.SearchDataSourcesOutput, error)
	SearchFolders(context.Context, *quicksight.SearchFoldersInput, ...func(*quicksight.Options)) (*quicksight.SearchFoldersOutput, error)
	SearchGroups(context.Context, *quicksight.SearchGroupsInput, ...func(*quicksight.Options)) (*quicksight.SearchGroupsOutput, error)
}
