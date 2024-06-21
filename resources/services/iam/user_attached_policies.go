package iam

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/cloudquery/resources/services/iam/models"
	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func userAttachedPolicies() *schema.Table {
	tableName := "aws_iam_user_attached_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html`,
		Resolver:    fetchIamUserAttachedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "user_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "policy_name",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("PolicyName"),
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

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*ohaws.User)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	config := iam.ListAttachedUserPoliciesInput{
		UserName: p.UserName,
	}
	paginator := iam.NewListAttachedUserPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
	}
	return nil
}

func resolveAccessKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	key := resource.Item.(models.AccessKeyWrapper)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "iam",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  "access-key/" + aws.ToString(key.UserName) + "/" + aws.ToString(key.AccessKeyId),
	}
	return resource.Set(c.Name, a.String())
}

func getUserAccessKey(ctx context.Context, svc services.IamClient, userName, keyID string) (*ohaws.AccessKey, error) {
	keysResp, err := svc.ListAccessKeys(ctx, &iam.ListAccessKeysInput{UserName: aws.String(userName)})
	if err != nil {
		return nil, fmt.Errorf("error getting user access keys: %w", err)
	}

	var k ohaws.AccessKey
	for _, md := range keysResp.AccessKeyMetadata {
		if md.AccessKeyId != nil && *md.AccessKeyId == keyID {
			k.AccessKeyMetadata = md
			break
		}
	}

	if k.AccessKeyId == nil {
		return nil, fmt.Errorf("key with id not found: %s", keyID)
	}

	r, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: k.AccessKeyId})
	if err != nil {
		return nil, fmt.Errorf("error getting access key last used: %w", err)
	}

	k.LastUsedDate = r.AccessKeyLastUsed.LastUsedDate
	if r.AccessKeyLastUsed.Region != nil {
		k.LastUsedRegion = *r.AccessKeyLastUsed.Region
	}
	if r.AccessKeyLastUsed.ServiceName != nil {
		k.LastUsedServiceName = *r.AccessKeyLastUsed.ServiceName
	}

	return &k, nil
}
