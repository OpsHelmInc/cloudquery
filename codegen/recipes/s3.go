package recipes

import (
	"reflect"
	"strings"

	"github.com/OpsHelmInc/ohaws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func S3Resources() []*Resource {
	resources := []*Resource{

		{
			SubService: "accounts",
			Struct:     &ohaws.PublicAccessBlockConfigurationWrapper{},
			SkipFields: []string{"ARN"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "buckets",
			Struct:     &ohaws.WrappedBucket{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveBucketARN()`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"BucketEncryptionRules()",
				"BucketLifecycles()",
				"BucketGrants()",
				"BucketCorsRules()",
			},
		},
		{
			SubService:  "bucket_encryption_rules",
			Struct:      &types.ServerSideEncryptionRule{},
			Description: "https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "sse_algorithm",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApplyServerSideEncryptionByDefault.SSEAlgorithm")`,
					},
				}...),
		},
		{
			SubService:  "bucket_lifecycles",
			Struct:      &types.LifecycleRule{},
			Description: "https://docs.aws.amazon.com/AmazonS3/latest/API/API_LifecycleRule.html",
			SkipFields:  []string{"Filter"},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "bucket_grants",
			Struct:      &types.Grant{},
			Description: "https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "bucket_cors_rules",
			Struct:      &types.CORSRule{},
			Description: "https://docs.aws.amazon.com/AmazonS3/latest/API/API_CORSRule.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "bucket_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "s3"
		r.Multiplex = "client.AccountMultiplex"
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
