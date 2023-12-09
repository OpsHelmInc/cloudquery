package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

func fetchAppstreamStackUserAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeUserStackAssociationsInput
	input.StackName = parent.Item.(types.Stack).Name
	input.MaxResults = aws.Int32(25)

	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeUserStackAssociations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.UserStackAssociations
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
