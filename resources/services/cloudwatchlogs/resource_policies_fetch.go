package cloudwatchlogs

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

func fetchCloudwatchlogsResourcePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudwatchlogs.DescribeResourcePoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	for {
		response, err := svc.DescribeResourcePolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ResourcePolicies
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
