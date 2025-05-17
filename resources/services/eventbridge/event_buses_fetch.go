package eventbridge

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchEventbridgeEventBuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListEventBusesInput
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListEventBuses(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.EventBuses
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getEventbridgeEventBus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	eb := resource.Item.(types.EventBus)

	resp, err := svc.DescribeEventBus(ctx, &eventbridge.DescribeEventBusInput{Name: eb.Name})
	if err != nil {
		return fmt.Errorf("error describing event bus: %w", err)
	}

	eventBus := &ohaws.EventBus{
		DescribeEventBusOutput: resp,
	}

	tags, err := getEventBridgeTags(ctx, meta, *eventBus.Arn)
	if err != nil {
		return err
	}

	eventBus.Tags = tags
	resource.Item = eventBus

	return nil
}

func fetchEventbridgeEventBusRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*ohaws.EventBus)
	input := eventbridge.ListRulesInput{
		EventBusName: p.Arn,
	}
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	for {
		response, err := svc.ListRules(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Rules
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getEventBusRule(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Eventbridge
	r := resource.Item.(types.Rule)

	resp, err := svc.DescribeRule(ctx, &eventbridge.DescribeRuleInput{
		Name:         r.Name,
		EventBusName: r.EventBusName,
	})
	if err != nil {
		return fmt.Errorf("error describing event bridge rule: %w", err)
	}

	rule := &ohaws.EventBridgeRule{
		DescribeRuleOutput: resp,
	}

	tags, err := getEventBridgeTags(ctx, meta, *rule.Arn)
	if err != nil {
		return err
	}
	rule.Tags = tags

	resource.Item = rule

	return nil
}

func getEventBridgeTags(ctx context.Context, meta schema.ClientMeta, resourceArn string) ([]types.Tag, error) {
	cl := meta.(*client.Client)
	svc := cl.Services().Eventbridge
	input := eventbridge.ListTagsForResourceInput{
		ResourceARN: &resourceArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return nil, err
	}
	return output.Tags, nil
}
