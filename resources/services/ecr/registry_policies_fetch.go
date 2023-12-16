package ecr

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEcrRegistryPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ecr
	output, err := svc.GetRegistryPolicy(ctx, &ecr.GetRegistryPolicyInput{})
	if err != nil {
		if client.IsAWSError(err, "RegistryPolicyNotFoundException") {
			return nil
		}
		return err
	}
	res <- output
	return nil
}
