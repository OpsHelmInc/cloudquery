package directconnect

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeVirtualInterfacesInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeVirtualInterfaces(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.VirtualInterfaces
	return nil
}

func resolveVirtualInterfaceARN() schema.ColumnResolver {
	return client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
		return []string{"dxvif", *resource.Item.(types.VirtualInterface).VirtualInterfaceId}, nil
	})
}
