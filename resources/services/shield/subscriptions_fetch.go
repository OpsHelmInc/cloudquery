package shield

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/shield"
)

func fetchShieldSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Shield
	config := shield.DescribeSubscriptionInput{}
	output, err := svc.DescribeSubscription(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- output.Subscription
	return nil
}
