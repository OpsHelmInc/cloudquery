package iam

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/iam/models"
)

func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.GetAccountPasswordPolicyInput
	c := meta.(*client.Client)
	svc := c.Services().Iam
	response, err := svc.GetAccountPasswordPolicy(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			res <- models.PasswordPolicyWrapper{PolicyExists: false}
			return nil
		}
		return err
	}
	res <- models.PasswordPolicyWrapper{PasswordPolicy: *response.PasswordPolicy, PolicyExists: true}
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
