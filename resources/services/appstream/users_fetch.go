package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

func fetchAppstreamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeUsersInput
	input.AuthenticationType = types.AuthenticationTypeUserpool
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeUsers(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Users
		if response.NextToken == nil {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
