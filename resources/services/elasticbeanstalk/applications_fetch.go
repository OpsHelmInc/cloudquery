package elasticbeanstalk

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
)

func fetchElasticbeanstalkApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeApplicationsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticbeanstalk
	output, err := svc.DescribeApplications(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.Applications
	return nil
}
