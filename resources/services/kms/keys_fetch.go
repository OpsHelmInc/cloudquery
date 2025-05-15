package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchKmsKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms

	config := kms.ListKeysInput{Limit: aws.Int32(1000)}
	p := kms.NewListKeysPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Keys
	}
	return nil
}

func getKey(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Kms
	item := resource.Item.(types.KeyListEntry)

	d, err := svc.DescribeKey(ctx, &kms.DescribeKeyInput{KeyId: item.KeyId})
	if err != nil {
		return err
	}

	var tags []ohaws.Tag
	if d.KeyMetadata.Origin != "EXTERNAL" && d.KeyMetadata.KeyManager != "AWS" {
		t, err := getKeyTags(ctx, c, svc, *d.KeyMetadata.KeyId)
		if err != nil {
			return err
		}
		tags = t
	}

	resource.Item = &ohaws.KMSKey{
		KeyMetadata: *d.KeyMetadata,
		Tags:        tags,
	}

	return nil
}

func resolveKeysReplicaKeys(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	key := resource.Item.(*ohaws.KMSKey)
	if key.MultiRegionConfiguration == nil {
		return nil
	}
	return resource.Set(c.Name, key.MultiRegionConfiguration.ReplicaKeys)
}

func getKeyTags(ctx context.Context, cl *client.Client, svc services.KmsClient, keyID string) ([]ohaws.Tag, error) {
	params := kms.ListResourceTagsInput{KeyId: aws.String(keyID)}
	paginator := kms.NewListResourceTagsPaginator(svc, &params)
	var tags []ohaws.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil, err
		}

		for _, t := range page.Tags {
			tags = append(tags, ohaws.Tag{Key: t.TagKey, Value: t.TagValue})
		}
	}
	return tags, nil
}

func resolveKeysRotationEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Kms
	key := resource.Item.(*ohaws.KMSKey)
	if key.Origin == "EXTERNAL" || key.KeyManager == "AWS" {
		return nil
	}
	result, err := svc.GetKeyRotationStatus(ctx, &kms.GetKeyRotationStatusInput{KeyId: key.KeyId})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, result.KeyRotationEnabled)
}
