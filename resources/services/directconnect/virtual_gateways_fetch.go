package directconnect

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

func fetchDirectconnectVirtualGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeVirtualGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualGateways(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.VirtualGateways
	return nil
}
