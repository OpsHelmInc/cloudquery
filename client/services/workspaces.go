// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
)

//go:generate mockgen -package=mocks -destination=../mocks/workspaces.go -source=workspaces.go WorkspacesClient
type WorkspacesClient interface {
	DescribeAccount(context.Context, *workspaces.DescribeAccountInput, ...func(*workspaces.Options)) (*workspaces.DescribeAccountOutput, error)
	DescribeAccountModifications(context.Context, *workspaces.DescribeAccountModificationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeAccountModificationsOutput, error)
	DescribeApplicationAssociations(context.Context, *workspaces.DescribeApplicationAssociationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeApplicationAssociationsOutput, error)
	DescribeApplications(context.Context, *workspaces.DescribeApplicationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeApplicationsOutput, error)
	DescribeBundleAssociations(context.Context, *workspaces.DescribeBundleAssociationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeBundleAssociationsOutput, error)
	DescribeClientBranding(context.Context, *workspaces.DescribeClientBrandingInput, ...func(*workspaces.Options)) (*workspaces.DescribeClientBrandingOutput, error)
	DescribeClientProperties(context.Context, *workspaces.DescribeClientPropertiesInput, ...func(*workspaces.Options)) (*workspaces.DescribeClientPropertiesOutput, error)
	DescribeConnectClientAddIns(context.Context, *workspaces.DescribeConnectClientAddInsInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectClientAddInsOutput, error)
	DescribeConnectionAliasPermissions(context.Context, *workspaces.DescribeConnectionAliasPermissionsInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasPermissionsOutput, error)
	DescribeConnectionAliases(context.Context, *workspaces.DescribeConnectionAliasesInput, ...func(*workspaces.Options)) (*workspaces.DescribeConnectionAliasesOutput, error)
	DescribeImageAssociations(context.Context, *workspaces.DescribeImageAssociationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeImageAssociationsOutput, error)
	DescribeIpGroups(context.Context, *workspaces.DescribeIpGroupsInput, ...func(*workspaces.Options)) (*workspaces.DescribeIpGroupsOutput, error)
	DescribeTags(context.Context, *workspaces.DescribeTagsInput, ...func(*workspaces.Options)) (*workspaces.DescribeTagsOutput, error)
	DescribeWorkspaceAssociations(context.Context, *workspaces.DescribeWorkspaceAssociationsInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceAssociationsOutput, error)
	DescribeWorkspaceBundles(context.Context, *workspaces.DescribeWorkspaceBundlesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceDirectories(context.Context, *workspaces.DescribeWorkspaceDirectoriesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceImagePermissions(context.Context, *workspaces.DescribeWorkspaceImagePermissionsInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error)
	DescribeWorkspaceImages(context.Context, *workspaces.DescribeWorkspaceImagesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceImagesOutput, error)
	DescribeWorkspaceSnapshots(context.Context, *workspaces.DescribeWorkspaceSnapshotsInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceSnapshotsOutput, error)
	DescribeWorkspaces(context.Context, *workspaces.DescribeWorkspacesInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesConnectionStatus(context.Context, *workspaces.DescribeWorkspacesConnectionStatusInput, ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	GetAccountLink(context.Context, *workspaces.GetAccountLinkInput, ...func(*workspaces.Options)) (*workspaces.GetAccountLinkOutput, error)
	ListAccountLinks(context.Context, *workspaces.ListAccountLinksInput, ...func(*workspaces.Options)) (*workspaces.ListAccountLinksOutput, error)
	ListAvailableManagementCidrRanges(context.Context, *workspaces.ListAvailableManagementCidrRangesInput, ...func(*workspaces.Options)) (*workspaces.ListAvailableManagementCidrRangesOutput, error)
}
