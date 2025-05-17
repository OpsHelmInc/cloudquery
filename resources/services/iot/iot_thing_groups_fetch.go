package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIotThingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingGroupsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	for {
		response, err := svc.ListThingGroups(ctx, &input)
		if err != nil {
			return err
		}
		for _, g := range response.ThingGroups {
			group, err := svc.DescribeThingGroup(ctx, &iot.DescribeThingGroupInput{
				ThingGroupName: g.GroupName,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- &ohaws.IoTThingGroup{
				DescribeThingGroupOutput: *group,
			}
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getIotThingGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	group := resource.Item.(*ohaws.IoTThingGroup)

	tags, err := getResourceTags(ctx, svc, aws.ToString(group.ThingGroupArn))
	if err != nil {
		return fmt.Errorf("error listing tags: %w", err)
	}

	group.Tags = tags
	resource.Item = group

	return nil
}

func ResolveIotThingGroupThingsInGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*ohaws.IoTThingGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingsInThingGroupInput{
		ThingGroupName: i.ThingGroupName,
		MaxResults:     aws.Int32(250),
	}

	var things []string
	for {
		response, err := svc.ListThingsInThingGroup(ctx, &input)
		if err != nil {
			return err
		}

		things = append(things, response.Things...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, things)
}

func ResolveIotThingGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*ohaws.IoTThingGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListAttachedPoliciesInput{
		Target:   i.ThingGroupArn,
		PageSize: aws.Int32(250),
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
