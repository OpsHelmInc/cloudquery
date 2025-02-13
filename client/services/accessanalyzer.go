// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

//go:generate mockgen -package=mocks -destination=../mocks/accessanalyzer.go -source=accessanalyzer.go AccessanalyzerClient
type AccessanalyzerClient interface {
	GetAccessPreview(context.Context, *accessanalyzer.GetAccessPreviewInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAccessPreviewOutput, error)
	GetAnalyzedResource(context.Context, *accessanalyzer.GetAnalyzedResourceInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzedResourceOutput, error)
	GetAnalyzer(context.Context, *accessanalyzer.GetAnalyzerInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetAnalyzerOutput, error)
	GetArchiveRule(context.Context, *accessanalyzer.GetArchiveRuleInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetArchiveRuleOutput, error)
	GetFinding(context.Context, *accessanalyzer.GetFindingInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingOutput, error)
	GetFindingRecommendation(context.Context, *accessanalyzer.GetFindingRecommendationInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingRecommendationOutput, error)
	GetFindingV2(context.Context, *accessanalyzer.GetFindingV2Input, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetFindingV2Output, error)
	GetGeneratedPolicy(context.Context, *accessanalyzer.GetGeneratedPolicyInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.GetGeneratedPolicyOutput, error)
	ListAccessPreviewFindings(context.Context, *accessanalyzer.ListAccessPreviewFindingsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewFindingsOutput, error)
	ListAccessPreviews(context.Context, *accessanalyzer.ListAccessPreviewsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAccessPreviewsOutput, error)
	ListAnalyzedResources(context.Context, *accessanalyzer.ListAnalyzedResourcesInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzedResourcesOutput, error)
	ListAnalyzers(context.Context, *accessanalyzer.ListAnalyzersInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error)
	ListArchiveRules(context.Context, *accessanalyzer.ListArchiveRulesInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListArchiveRulesOutput, error)
	ListFindings(context.Context, *accessanalyzer.ListFindingsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsOutput, error)
	ListFindingsV2(context.Context, *accessanalyzer.ListFindingsV2Input, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsV2Output, error)
	ListPolicyGenerations(context.Context, *accessanalyzer.ListPolicyGenerationsInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListPolicyGenerationsOutput, error)
	ListTagsForResource(context.Context, *accessanalyzer.ListTagsForResourceInput, ...func(*accessanalyzer.Options)) (*accessanalyzer.ListTagsForResourceOutput, error)
}
