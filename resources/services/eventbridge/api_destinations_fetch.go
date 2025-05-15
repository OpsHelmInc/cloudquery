package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchEventbridgeApiDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListApiDestinationsInput
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListApiDestinations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ApiDestinations
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getEventBridgeAPIDestination(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	d := resource.Item.(types.ApiDestination)

	resp, err := svc.DescribeApiDestination(ctx, &eventbridge.DescribeApiDestinationInput{Name: d.Name})
	if err != nil {
		return err
	}

	dest := &ohaws.EventBridgeAPIDestination{
		DescribeApiDestinationOutput: resp,
	}
	resource.Item = dest

	return nil
}
