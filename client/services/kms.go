// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

//go:generate mockgen -package=mocks -destination=../mocks/kms.go -source=kms.go KmsClient
type KmsClient interface {
	DescribeCustomKeyStores(context.Context, *kms.DescribeCustomKeyStoresInput, ...func(*kms.Options)) (*kms.DescribeCustomKeyStoresOutput, error)
	DescribeKey(context.Context, *kms.DescribeKeyInput, ...func(*kms.Options)) (*kms.DescribeKeyOutput, error)
	GetKeyPolicy(context.Context, *kms.GetKeyPolicyInput, ...func(*kms.Options)) (*kms.GetKeyPolicyOutput, error)
	GetKeyRotationStatus(context.Context, *kms.GetKeyRotationStatusInput, ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error)
	GetParametersForImport(context.Context, *kms.GetParametersForImportInput, ...func(*kms.Options)) (*kms.GetParametersForImportOutput, error)
	GetPublicKey(context.Context, *kms.GetPublicKeyInput, ...func(*kms.Options)) (*kms.GetPublicKeyOutput, error)
	ListAliases(context.Context, *kms.ListAliasesInput, ...func(*kms.Options)) (*kms.ListAliasesOutput, error)
	ListGrants(context.Context, *kms.ListGrantsInput, ...func(*kms.Options)) (*kms.ListGrantsOutput, error)
	ListKeyPolicies(context.Context, *kms.ListKeyPoliciesInput, ...func(*kms.Options)) (*kms.ListKeyPoliciesOutput, error)
	ListKeyRotations(context.Context, *kms.ListKeyRotationsInput, ...func(*kms.Options)) (*kms.ListKeyRotationsOutput, error)
	ListKeys(context.Context, *kms.ListKeysInput, ...func(*kms.Options)) (*kms.ListKeysOutput, error)
	ListResourceTags(context.Context, *kms.ListResourceTagsInput, ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error)
	ListRetirableGrants(context.Context, *kms.ListRetirableGrantsInput, ...func(*kms.Options)) (*kms.ListRetirableGrantsOutput, error)
}
