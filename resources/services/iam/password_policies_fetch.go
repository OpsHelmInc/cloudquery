package iam

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/iam/models"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/plugin-sdk/schema"
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
