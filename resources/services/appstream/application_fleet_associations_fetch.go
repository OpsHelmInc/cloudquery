package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

func fetchAppstreamApplicationFleetAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	parentApplication := parent.Item.(types.Application)

	var input appstream.DescribeApplicationFleetAssociationsInput
	input.ApplicationArn = parentApplication.Arn

	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeApplicationFleetAssociations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ApplicationFleetAssociations
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
