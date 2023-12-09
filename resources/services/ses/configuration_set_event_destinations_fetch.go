package ses

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func fetchSesConfigurationSetEventDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	s := parent.Item.(*sesv2.GetConfigurationSetOutput)

	output, err := svc.GetConfigurationSetEventDestinations(ctx, &sesv2.GetConfigurationSetEventDestinationsInput{
		ConfigurationSetName: s.ConfigurationSetName,
	})
	if err != nil {
		return err
	}
	res <- output.EventDestinations
	return nil
}
