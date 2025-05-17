package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.GetAccountPasswordPolicyInput
	c := meta.(*client.Client)
	svc := c.Services().Iam
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			res <- ohaws.PasswordPolicy{PolicyExists: false}
			return nil
		}
		return err
	}
	res <- ohaws.PasswordPolicy{PasswordPolicy: *response.PasswordPolicy, PolicyExists: true}
	return nil
}

func resolvePasswordPolicyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "iam",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  "password-policy",
	}
	return resource.Set(c.Name, a.String())
}
