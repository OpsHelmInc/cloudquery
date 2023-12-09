package directconnect

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

func fetchDirectconnectLags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config directconnect.DescribeLagsInput
	c := meta.(*client.Client)
	svc := c.Services().Directconnect
	output, err := svc.DescribeLags(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Lags
	return nil
}

func resolveLagARN() schema.ColumnResolver {
	return client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
		return []string{"dxlag", *resource.Item.(types.Lag).LagId}, nil
	})
}
