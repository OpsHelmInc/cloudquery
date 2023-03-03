package ec2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2LaunchTemplateVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := ec2.DescribeLaunchTemplateVersionsInput{
		Versions:   []string{"$Latest", "$Default"},
		MaxResults: aws.Int32(200),
	}
	c := meta.(*client.Client)
	svc := c.Services().Ec2

	for {
		output, err := svc.DescribeLaunchTemplateVersions(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.LaunchTemplateVersions
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
