package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIotThings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	for {
		response, err := svc.ListThings(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Things
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getIotThing(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	i := resource.Item.(types.ThingAttribute)

	input := iot.DescribeThingInput{
		ThingName: i.ThingName,
	}
	resp, err := svc.DescribeThing(ctx, &input)
	if err != nil {
		return fmt.Errorf("error describing IoT thing: %w", err)
	}

	t := ohaws.IoTThing{
		DescribeThingOutput: *resp,
	}

	tags, err := getResourceTags(ctx, svc, aws.ToString(t.ThingArn))
	if err != nil {
		return fmt.Errorf("error listing tags: %w", err)
	}

	t.Tags = tags
	resource.Item = &t

	return nil
}

func ResolveIotThingPrincipals(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*ohaws.IoTThing)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingPrincipalsInput{
		ThingName:  i.ThingName,
		MaxResults: aws.Int32(250),
	}
	var principals []string

	for {
		response, err := svc.ListThingPrincipals(ctx, &input)
		if err != nil {
			return err
		}
		principals = append(principals, response.Principals...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, principals)
}

func getResourceTags(ctx context.Context, svc services.IotClient, resourceArn string) ([]types.Tag, error) {
	var tags []types.Tag

	input := iot.ListTagsForResourceInput{
		ResourceArn: aws.String(resourceArn),
	}

	for {
		response, err := svc.ListTagsForResource(ctx, &input)
		if err != nil {
			return nil, err
		}

		tags = append(tags, response.Tags...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return tags, nil
}
