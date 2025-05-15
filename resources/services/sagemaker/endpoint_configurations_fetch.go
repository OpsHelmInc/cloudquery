package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchSagemakerEndpointConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListEndpointConfigsInput{}
	for {
		response, err := svc.ListEndpointConfigs(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.EndpointConfigs

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getEndpointConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.EndpointConfigSummary)

	response, err := svc.DescribeEndpointConfig(ctx, &sagemaker.DescribeEndpointConfigInput{
		EndpointConfigName: n.EndpointConfigName,
	})
	if err != nil {
		return err
	}

	var ec ohaws.SagemakerEndpointConfig
	if err := mapstructure.Decode(response, &ec); err != nil {
		return err
	}

	listTagsInput := &sagemaker.ListTagsInput{
		ResourceArn: ec.EndpointConfigArn,
	}

	tagsPaginator := sagemaker.NewListTagsPaginator(svc, listTagsInput)
	for tagsPaginator.HasMorePages() {
		p, err := tagsPaginator.NextPage(ctx, func(o *sagemaker.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}

		ec.Tags = append(ec.Tags, p.Tags...)
	}

	resource.Item = &ec
	return nil
}
