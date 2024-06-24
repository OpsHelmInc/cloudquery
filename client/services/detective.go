// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/detective"
)

//go:generate mockgen -package=mocks -destination=../mocks/detective.go -source=detective.go DetectiveClient
type DetectiveClient interface {
	BatchGetGraphMemberDatasources(context.Context, *detective.BatchGetGraphMemberDatasourcesInput, ...func(*detective.Options)) (*detective.BatchGetGraphMemberDatasourcesOutput, error)
	BatchGetMembershipDatasources(context.Context, *detective.BatchGetMembershipDatasourcesInput, ...func(*detective.Options)) (*detective.BatchGetMembershipDatasourcesOutput, error)
	DescribeOrganizationConfiguration(context.Context, *detective.DescribeOrganizationConfigurationInput, ...func(*detective.Options)) (*detective.DescribeOrganizationConfigurationOutput, error)
	GetInvestigation(context.Context, *detective.GetInvestigationInput, ...func(*detective.Options)) (*detective.GetInvestigationOutput, error)
	GetMembers(context.Context, *detective.GetMembersInput, ...func(*detective.Options)) (*detective.GetMembersOutput, error)
	ListDatasourcePackages(context.Context, *detective.ListDatasourcePackagesInput, ...func(*detective.Options)) (*detective.ListDatasourcePackagesOutput, error)
	ListGraphs(context.Context, *detective.ListGraphsInput, ...func(*detective.Options)) (*detective.ListGraphsOutput, error)
	ListIndicators(context.Context, *detective.ListIndicatorsInput, ...func(*detective.Options)) (*detective.ListIndicatorsOutput, error)
	ListInvestigations(context.Context, *detective.ListInvestigationsInput, ...func(*detective.Options)) (*detective.ListInvestigationsOutput, error)
	ListInvitations(context.Context, *detective.ListInvitationsInput, ...func(*detective.Options)) (*detective.ListInvitationsOutput, error)
	ListMembers(context.Context, *detective.ListMembersInput, ...func(*detective.Options)) (*detective.ListMembersOutput, error)
	ListOrganizationAdminAccounts(context.Context, *detective.ListOrganizationAdminAccountsInput, ...func(*detective.Options)) (*detective.ListOrganizationAdminAccountsOutput, error)
	ListTagsForResource(context.Context, *detective.ListTagsForResourceInput, ...func(*detective.Options)) (*detective.ListTagsForResourceOutput, error)
}
