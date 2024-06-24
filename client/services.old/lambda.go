// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

//go:generate mockgen -package=mocks -destination=../mocks/lambda.go -source=lambda.go LambdaClient
type LambdaClient interface {
	GetAccountSettings(context.Context, *lambda.GetAccountSettingsInput, ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error)
	GetAlias(context.Context, *lambda.GetAliasInput, ...func(*lambda.Options)) (*lambda.GetAliasOutput, error)
	GetCodeSigningConfig(context.Context, *lambda.GetCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	GetEventSourceMapping(context.Context, *lambda.GetEventSourceMappingInput, ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error)
	GetFunction(context.Context, *lambda.GetFunctionInput, ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	GetFunctionCodeSigningConfig(context.Context, *lambda.GetFunctionCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionConcurrency(context.Context, *lambda.GetFunctionConcurrencyInput, ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConfiguration(context.Context, *lambda.GetFunctionConfigurationInput, ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error)
	GetFunctionEventInvokeConfig(context.Context, *lambda.GetFunctionEventInvokeConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionUrlConfig(context.Context, *lambda.GetFunctionUrlConfigInput, ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error)
	GetLayerVersion(context.Context, *lambda.GetLayerVersionInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionByArn(context.Context, *lambda.GetLayerVersionByArnInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionPolicy(context.Context, *lambda.GetLayerVersionPolicyInput, ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	GetPolicy(context.Context, *lambda.GetPolicyInput, ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	GetProvisionedConcurrencyConfig(context.Context, *lambda.GetProvisionedConcurrencyConfigInput, ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetRuntimeManagementConfig(context.Context, *lambda.GetRuntimeManagementConfigInput, ...func(*lambda.Options)) (*lambda.GetRuntimeManagementConfigOutput, error)
	ListAliases(context.Context, *lambda.ListAliasesInput, ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	ListCodeSigningConfigs(context.Context, *lambda.ListCodeSigningConfigsInput, ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error)
	ListEventSourceMappings(context.Context, *lambda.ListEventSourceMappingsInput, ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	ListFunctionEventInvokeConfigs(context.Context, *lambda.ListFunctionEventInvokeConfigsInput, ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionUrlConfigs(context.Context, *lambda.ListFunctionUrlConfigsInput, ...func(*lambda.Options)) (*lambda.ListFunctionUrlConfigsOutput, error)
	ListFunctions(context.Context, *lambda.ListFunctionsInput, ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	ListFunctionsByCodeSigningConfig(context.Context, *lambda.ListFunctionsByCodeSigningConfigInput, ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListLayerVersions(context.Context, *lambda.ListLayerVersionsInput, ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	ListLayers(context.Context, *lambda.ListLayersInput, ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	ListProvisionedConcurrencyConfigs(context.Context, *lambda.ListProvisionedConcurrencyConfigsInput, ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListTags(context.Context, *lambda.ListTagsInput, ...func(*lambda.Options)) (*lambda.ListTagsOutput, error)
	ListVersionsByFunction(context.Context, *lambda.ListVersionsByFunctionInput, ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
}