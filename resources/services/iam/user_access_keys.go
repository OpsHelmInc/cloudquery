package iam

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/ohaws"
)

func userAccessKeys() *schema.Table {
	tableName := "aws_iam_user_access_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html`,
		Resolver:    fetchIamUserAccessKeys,
		// PostResourceResolver: postIamUserAccessKeyResolver,
		Transform: transformers.TransformWithStruct(&ohaws.AccessKey{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "user_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "access_key_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("AccessKeyId"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
		},
	}
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(*ohaws.User)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	config.UserName = p.UserName
	paginator := iam.NewListAccessKeysPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		keys := make([]ohaws.AccessKey, len(page.AccessKeyMetadata))
		for i, key := range page.AccessKeyMetadata {
			keys[i] = ohaws.AccessKey{AccessKeyMetadata: key}

			r, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: keys[i].AccessKeyId}, func(o *iam.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return fmt.Errorf("error getting access key last used: %w", err)
			}

			keys[i].LastUsedDate = r.AccessKeyLastUsed.LastUsedDate
			keys[i].LastUsedRegion = aws.ToString(r.AccessKeyLastUsed.Region)
			keys[i].LastUsedServiceName = aws.ToString(r.AccessKeyLastUsed.ServiceName)
		}
		res <- keys
	}
	return nil
}
