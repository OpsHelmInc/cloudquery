package elasticbeanstalk

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func fetchElasticbeanstalkApplicationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeApplicationVersionsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticbeanstalk

	for {
		output, err := svc.DescribeApplicationVersions(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ApplicationVersions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
