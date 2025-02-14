package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func userPolicies() *schema.Table {
	tableName := "aws_iam_user_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetUserPolicy.html`,
		Resolver:            fetchIamUserPolicies,
		PreResourceResolver: getUserPolicy,
		Transform:           transformers.TransformWithStruct(&iam.GetUserPolicyOutput{}, transformers.WithPrimaryKeyComponents("PolicyName"), transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "user_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
			{
				Name:     "policy_document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamUserPolicyPolicyDocument,
			},
		},
	}
}

func fetchIamUserPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	user := parent.Item.(*ohaws.User)
	config := iam.ListUserPoliciesInput{UserName: user.UserName}
	paginator := iam.NewListUserPoliciesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.PolicyNames
	}
	return nil
}

func getUserPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	p := resource.Item.(string)
	user := resource.Parent.Item.(*ohaws.User)

	policyResult, err := svc.GetUserPolicy(ctx, &iam.GetUserPolicyInput{PolicyName: &p, UserName: user.UserName}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = policyResult
	return nil
}

func resolveIamUserPolicyPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetUserPolicyOutput)

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return err
	}

	var document map[string]any
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, document)
}
