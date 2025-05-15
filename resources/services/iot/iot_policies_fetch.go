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

func fetchIotPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListPoliciesInput{
		PageSize: aws.Int32(250),
	}

	for {
		response, err := svc.ListPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, s := range response.Policies {
			profile, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
				PolicyName: s.PolicyName,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- &ohaws.IoTPolicy{
				GetPolicyOutput: *profile,
			}
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}

func getIotPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	policy := resource.Item.(*ohaws.IoTPolicy)

	tags, err := getResourceTags(ctx, svc, aws.ToString(policy.PolicyArn))
	if err != nil {
		return fmt.Errorf("error listing tags: %w", err)
	}

	policy.Tags = tags
	resource.Item = policy

	return nil
}

func ResolveIotPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*ohaws.IoTPolicy)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.PolicyArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input)
		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
