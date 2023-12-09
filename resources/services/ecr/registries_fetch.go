package ecr

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func fetchEcrRegistries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ecr
	output, err := svc.DescribeRegistry(ctx, &ecr.DescribeRegistryInput{})
	if err != nil {
		return err
	}
	res <- output
	return nil
}
