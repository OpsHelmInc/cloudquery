package ec2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2LaunchTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := ec2.DescribeLaunchTemplatesInput{
		MaxResults: aws.Int32(200),
	}
	c := meta.(*client.Client)
	svc := c.Services().Ec2

	for {
		output, err := svc.DescribeLaunchTemplates(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.LaunchTemplates
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getLaunchTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	launchTemplate := resource.Item.(types.LaunchTemplate)

	output, err := svc.GetLaunchTemplateData(ctx, &ec2.GetLaunchTemplateDataInput{InstanceId: launchTemplate.LaunchTemplateId})
	if err != nil {
		return err
	}

	resource.Item = models.LaunchTemplateWrapper{
		LaunchTemplateId:   launchTemplate.LaunchTemplateId,
		LaunchTemplateData: output.LaunchTemplateData,
	}
	return nil
}
