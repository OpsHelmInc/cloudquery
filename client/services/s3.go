// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//go:generate mockgen -package=mocks -destination=../mocks/s3.go -source=s3.go S3Client
type S3Client interface {
	GetBucketAccelerateConfiguration(context.Context, *s3.GetBucketAccelerateConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error)
	GetBucketAcl(context.Context, *s3.GetBucketAclInput, ...func(*s3.Options)) (*s3.GetBucketAclOutput, error)
	GetBucketAnalyticsConfiguration(context.Context, *s3.GetBucketAnalyticsConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error)
	GetBucketCors(context.Context, *s3.GetBucketCorsInput, ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error)
	GetBucketEncryption(context.Context, *s3.GetBucketEncryptionInput, ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error)
	GetBucketIntelligentTieringConfiguration(context.Context, *s3.GetBucketIntelligentTieringConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error)
	GetBucketInventoryConfiguration(context.Context, *s3.GetBucketInventoryConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error)
	GetBucketLifecycleConfiguration(context.Context, *s3.GetBucketLifecycleConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error)
	GetBucketLocation(context.Context, *s3.GetBucketLocationInput, ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error)
	GetBucketLogging(context.Context, *s3.GetBucketLoggingInput, ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error)
	GetBucketMetadataTableConfiguration(context.Context, *s3.GetBucketMetadataTableConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketMetadataTableConfigurationOutput, error)
	GetBucketMetricsConfiguration(context.Context, *s3.GetBucketMetricsConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error)
	GetBucketNotificationConfiguration(context.Context, *s3.GetBucketNotificationConfigurationInput, ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error)
	GetBucketOwnershipControls(context.Context, *s3.GetBucketOwnershipControlsInput, ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error)
	GetBucketPolicy(context.Context, *s3.GetBucketPolicyInput, ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error)
	GetBucketPolicyStatus(context.Context, *s3.GetBucketPolicyStatusInput, ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error)
	GetBucketReplication(context.Context, *s3.GetBucketReplicationInput, ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error)
	GetBucketRequestPayment(context.Context, *s3.GetBucketRequestPaymentInput, ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error)
	GetBucketTagging(context.Context, *s3.GetBucketTaggingInput, ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error)
	GetBucketVersioning(context.Context, *s3.GetBucketVersioningInput, ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error)
	GetBucketWebsite(context.Context, *s3.GetBucketWebsiteInput, ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error)
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectAcl(context.Context, *s3.GetObjectAclInput, ...func(*s3.Options)) (*s3.GetObjectAclOutput, error)
	GetObjectAttributes(context.Context, *s3.GetObjectAttributesInput, ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
	GetObjectLegalHold(context.Context, *s3.GetObjectLegalHoldInput, ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error)
	GetObjectLockConfiguration(context.Context, *s3.GetObjectLockConfigurationInput, ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error)
	GetObjectRetention(context.Context, *s3.GetObjectRetentionInput, ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error)
	GetObjectTagging(context.Context, *s3.GetObjectTaggingInput, ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error)
	GetObjectTorrent(context.Context, *s3.GetObjectTorrentInput, ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error)
	GetPublicAccessBlock(context.Context, *s3.GetPublicAccessBlockInput, ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error)
	ListBucketAnalyticsConfigurations(context.Context, *s3.ListBucketAnalyticsConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error)
	ListBucketIntelligentTieringConfigurations(context.Context, *s3.ListBucketIntelligentTieringConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error)
	ListBucketInventoryConfigurations(context.Context, *s3.ListBucketInventoryConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error)
	ListBucketMetricsConfigurations(context.Context, *s3.ListBucketMetricsConfigurationsInput, ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error)
	ListBuckets(context.Context, *s3.ListBucketsInput, ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
	ListDirectoryBuckets(context.Context, *s3.ListDirectoryBucketsInput, ...func(*s3.Options)) (*s3.ListDirectoryBucketsOutput, error)
	ListMultipartUploads(context.Context, *s3.ListMultipartUploadsInput, ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error)
	ListObjectVersions(context.Context, *s3.ListObjectVersionsInput, ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error)
	ListObjects(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) (*s3.ListObjectsOutput, error)
	ListObjectsV2(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	ListParts(context.Context, *s3.ListPartsInput, ...func(*s3.Options)) (*s3.ListPartsOutput, error)
}
