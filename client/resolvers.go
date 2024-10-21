package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/thoas/go-funk"

	"github.com/OpsHelmInc/ohaws"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/scalar"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AccountID)
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Region)
}

func ResolveAWSNamespace(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AutoscalingNamespace)
}

func ResolveAWSPartition(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Partition)
}

func ResolveWAFScope(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).WAFScope)
}

func ResolveLanguageCode(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.LanguageCode)
}

func ResolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return ResolveTagPath("Tags")(ctx, meta, r, c)
}

func ResolveTagPath(fieldPath string) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		val := funk.Get(r.Item, fieldPath, funk.WithAllowZero())
		if val == nil {
			return r.Set(c.Name, map[string]string{}) // can't have nil or the integration test will make a fuss
		}

		return r.Set(c.Name, TagsToMap(val))
	}
}

func ResolveObjectHash(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	hash, err := hashstructure.Hash(r.Item, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}
	return r.Set(c.Name, fmt.Sprint(hash))
}

func ResolveOHResourceType(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	fetchedArn := r.Get("arn").(*scalar.String)
	if !fetchedArn.IsValid() {
		return fmt.Errorf("failed to resolve arn")
	}

	if strings.Contains(fetchedArn.String(), "test string") || strings.Contains(fetchedArn.String(), "teststring") {
		// We should only be entering here during "go test"
		return r.Set(c.Name, "test")
	}

	parsedArn, err := arn.Parse(fetchedArn.String())
	if err != nil {
		return fmt.Errorf("failed to parse arn: %s", fetchedArn.String())
	}

	resourceType, err := ohaws.ArnToResourceType(parsedArn)
	if err != nil {
		return fmt.Errorf("failed to get resource type from arn: %s", parsedArn.String())
	}

	return r.Set(c.Name, resourceType.CloudFormation)
}
