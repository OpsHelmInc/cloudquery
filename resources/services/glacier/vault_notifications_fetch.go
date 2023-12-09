package glacier

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

func fetchGlacierVaultNotifications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier
	p := parent.Item.(types.DescribeVaultOutput)

	response, err := svc.GetVaultNotifications(ctx, &glacier.GetVaultNotificationsInput{
		VaultName: p.VaultName,
	})
	if err != nil {
		return err
	}
	res <- response.VaultNotificationConfig
	return nil
}
