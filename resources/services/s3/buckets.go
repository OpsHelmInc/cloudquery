package s3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Buckets() *schema.Table {
	tableName := "aws_s3_buckets"
	return &schema.Table{
		Name:                "aws_s3_buckets",
		Resolver:            listS3Buckets,
		PreResourceResolver: resolveS3BucketsAttributes,
		Description:         `https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html`,
		Transform:           transformers.TransformWithStruct(&ohaws.WrappedBucket{}),
		Multiplex:           client.AccountMultiplex(tableName),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveBucketARN(),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			bucketCorsRules(),
			bucketEncryptionRules(),
			bucketGrants(),
			bucketLifecycles(),
			bucketNotificationConfigurations(),
			bucketObjectLockConfigurations(),
			bucketWebsites(),
			bucketLogging(),
			bucketOwnershipControls(),
			bucketReplications(),
			bucketPublicAccessBlock(),
			bucketVersionings(),
			bucketPolicies(),
		},
	}
}

func listS3Buckets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	response, err := svc.ListBuckets(ctx, nil, func(o *s3.Options) {
		o.Region = listBucketRegion(cl)
	})
	if err != nil {
		return err
	}
	for _, bucket := range response.Buckets {
		res <- &ohaws.WrappedBucket{
			Name:         bucket.Name,
			CreationDate: bucket.CreationDate,
		}
	}
	return nil
}

// listBucketRegion identifies the canonical region for S3 based on the partition
// in the future we might want to make this configurable if users are alright with the fact that performing this
// action in different regions will return different results
func listBucketRegion(cl *client.Client) string {
	switch cl.Partition {
	case "aws-cn":
		return "cn-north-1"
	case "aws-us-gov":
		return "us-gov-west-1"
	default:
		return "us-east-1"
	}
}

func resolveS3BucketsAttributes(ctx context.Context, meta schema.ClientMeta, r *schema.Resource) error {
	resource := r.Item.(*ohaws.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3

	output, err := svc.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: resource.Name,
	}, func(o *s3.Options) {
		o.Region = listBucketRegion(cl)
	})
	if err != nil {
		if isBucketNotFoundError(cl, err) {
			return nil
		}
		return err
	}
	// AWS does not specify a region if bucket is in us-east-1, so as long as no error we can assume an empty string is us-east-1
	resource.Region = "us-east-1"
	if output != nil && output.LocationConstraint != "" {
		resource.Region = string(output.LocationConstraint)
	}
	if output != nil && output.LocationConstraint == "EU" {
		resource.Region = "eu-west-1"
	}
	var errAll []error

	resolvers := []func(context.Context, schema.ClientMeta, *ohaws.WrappedBucket) error{
		resolveBucketLogging,
		resolveBucketPolicy,
		resolveBucketPolicyStatus,
		resolveBucketVersioning,
		resolveBucketPublicAccessBlock,
		resolveBucketReplication,
		resolveBucketTagging,
		resolveBucketOwnershipControls,
		resolveBucketLifecycles,
		resolveIntelligentTieringConfigurations,
	}
	for _, resolver := range resolvers {
		if err := resolver(ctx, meta, resource); err != nil {
			// If we received any error other than NoSuchBucketError, we return as this indicates that the bucket has been deleted
			// and therefore no other attributes can be resolved
			if isBucketNotFoundError(cl, err) {
				r.Item = resource
				return errors.Join(errAll...)
			}
			// This enables 403 errors to be recorded, but not block subsequent resolver calls
			errAll = append(errAll, err)
		}
	}
	r.Item = resource
	return errors.Join(errAll...)
}

func resolveBucketLogging(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	svc := meta.(*client.Client).Services().S3
	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if loggingOutput.LoggingEnabled == nil {
		return nil
	}
	resource.LoggingTargetBucket = loggingOutput.LoggingEnabled.TargetBucket
	resource.LoggingTargetPrefix = loggingOutput.LoggingEnabled.TargetPrefix
	return nil
}

func resolveBucketPolicy(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	policyOutput, err := svc.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	// check if we got an error but its access denied we can continue
	if err != nil {
		// if we got an error, and it's not a NoSuchBucketError, return err
		if client.IsAWSError(err, "NoSuchBucketPolicy") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if policyOutput == nil || policyOutput.Policy == nil {
		return nil
	}
	var p ohaws.PolicyDocument
	err = json.Unmarshal([]byte(*policyOutput.Policy), &p)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON policy: %v", err)
	}
	resource.Policy = &p
	return nil
}

func resolveBucketPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	policyStatusOutput, err := svc.GetBucketPolicyStatus(ctx, &s3.GetBucketPolicyStatusInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	// check if we got an error but its access denied we can continue
	if err != nil {
		// if we got an error, and it's not a NoSuchBucketError, return err
		if client.IsAWSError(err, "NoSuchBucketPolicy") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if policyStatusOutput == nil {
		return nil
	}
	if policyStatusOutput.PolicyStatus != nil {
		resource.IsPublic = *policyStatusOutput.PolicyStatus.IsPublic
	}
	return nil
}

func resolveBucketVersioning(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	resource.VersioningStatus = versioningOutput.Status
	resource.VersioningMfaDelete = versioningOutput.MFADelete
	return nil
}

func resolveBucketPublicAccessBlock(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	publicAccessOutput, err := svc.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		if isBucketNotFoundError(c, err) {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	resource.BlockPublicAcls = *publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicAcls
	resource.BlockPublicPolicy = *publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicPolicy
	resource.IgnorePublicAcls = *publicAccessOutput.PublicAccessBlockConfiguration.IgnorePublicAcls
	resource.RestrictPublicBuckets = *publicAccessOutput.PublicAccessBlockConfiguration.RestrictPublicBuckets
	return nil
}

func resolveBucketReplication(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	replicationOutput, err := svc.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})

	if err != nil {
		// If we received any error other than ReplicationConfigurationNotFoundError, we return and error
		if client.IsAWSError(err, "ReplicationConfigurationNotFoundError") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if replicationOutput.ReplicationConfiguration == nil {
		return nil
	}
	resource.ReplicationRole = replicationOutput.ReplicationConfiguration.Role
	resource.ReplicationRules = replicationOutput.ReplicationConfiguration.Rules
	return nil
}

func resolveBucketTagging(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	taggingOutput, err := svc.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	if err != nil {
		// If buckets tags are not set it will return an error instead of empty result
		if client.IsAWSError(err, "NoSuchTagSet") {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	if taggingOutput == nil {
		return nil
	}
	tags := make(map[string]*string, len(taggingOutput.TagSet))
	for _, t := range taggingOutput.TagSet {
		tags[*t.Key] = t.Value
	}
	resource.Tags = tags
	return nil
}

func resolveBucketOwnershipControls(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3

	getBucketOwnershipControlOutput, err := svc.GetBucketOwnershipControls(ctx, &s3.GetBucketOwnershipControlsInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})

	if err != nil {
		// If buckets ownership controls are not set it will return an error instead of empty result
		if client.IsAWSError(err, "OwnershipControlsNotFoundError") {
			return nil
		}

		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}

		return err
	}

	if getBucketOwnershipControlOutput == nil {
		return nil
	}

	ownershipControlRules := getBucketOwnershipControlOutput.OwnershipControls.Rules

	if len(ownershipControlRules) == 0 {
		return nil
	}

	stringArray := make([]string, 0, len(ownershipControlRules))

	for _, ownershipControlRule := range ownershipControlRules {
		stringArray = append(stringArray, string(ownershipControlRule.ObjectOwnership))
	}

	resource.OwnershipControls = stringArray
	return nil
}

func resolveBucketLifecycles(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	lifecycleOutput, err := svc.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = resource.Region
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchLifecycleConfiguration") {
			return nil
		}
		return err
	}
	resource.LifecycleRules = lifecycleOutput.Rules
	return nil
}

func resolveIntelligentTieringConfigurations(ctx context.Context, meta schema.ClientMeta, resource *ohaws.WrappedBucket) error {
	c := meta.(*client.Client)
	svc := c.Services().S3
	configList := []types.IntelligentTieringConfiguration{}
	params := &s3.ListBucketIntelligentTieringConfigurationsInput{Bucket: resource.Name}
	for {
		tieringOutput, err := svc.ListBucketIntelligentTieringConfigurations(ctx, params, func(options *s3.Options) {
			options.Region = resource.Region
		})

		if err != nil {
			return err
		}

		configList = append(configList, tieringOutput.IntelligentTieringConfigurationList...)
		if tieringOutput.NextContinuationToken == nil {
			break
		}

		params.ContinuationToken = tieringOutput.NextContinuationToken
	}

	resource.IntelligentTieringConfigurations = configList
	return nil
}

func isBucketNotFoundError(cl *client.Client, err error) bool {
	if cl.IsNotFoundError(err) {
		return true
	}
	if err.Error() == "bucket not found" {
		return true
	}
	return false
}

func resolveBucketARN() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.S3Service, func(resource *schema.Resource) ([]string, error) {
		return []string{*resource.Item.(*ohaws.WrappedBucket).Name}, nil
	})
}
