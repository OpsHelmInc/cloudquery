package glacier

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlacierVaultAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier
	p := parent.Item.(types.DescribeVaultOutput)

	response, err := svc.GetVaultAccessPolicy(ctx, &glacier.GetVaultAccessPolicyInput{
		VaultName: p.VaultName,
	})
	if err != nil {
		return err
	}
	res <- response.Policy
	return nil
}
