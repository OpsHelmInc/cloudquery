// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAppstreamImageBuilders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input appstream.DescribeImageBuildersInput
	c := meta.(*client.Client)
	svc := c.Services().Appstream
	for {
		response, err := svc.DescribeImageBuilders(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ImageBuilders

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
