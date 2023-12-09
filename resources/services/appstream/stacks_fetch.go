package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
)

func fetchAppstreamStacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeStacksInput
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeStacks(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Stacks
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
