package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

func fetchAppstreamStackEntitlements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeEntitlementsInput
	input.StackName = parent.Item.(types.Stack).Name
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeEntitlements(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Entitlements
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
