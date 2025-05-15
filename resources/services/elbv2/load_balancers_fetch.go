package elbv2

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElbv2LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elbv2.DescribeLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config)
		if err != nil {
			return err
		}
		lbs := make([]*ohaws.LoadBalancerV2, len(response.LoadBalancers))
		for idx, lb := range response.LoadBalancers {
			lbs[idx] = &ohaws.LoadBalancerV2{
				LoadBalancer: lb,
			}
		}
		res <- lbs
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func resolveElbv2loadBalancerWebACLArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(*ohaws.LoadBalancerV2)
	// only application load balancer can have web acl arn
	if p.Type != types.LoadBalancerTypeEnumApplication {
		return nil
	}
	cl := meta.(*client.Client).Services().Wafv2
	input := wafv2.GetWebACLForResourceInput{ResourceArn: p.LoadBalancerArn}
	response, err := cl.GetWebACLForResource(ctx, &input, func(options *wafv2.Options) {})
	if err != nil {
		var exc *wafv2types.WAFNonexistentItemException
		if errors.As(err, &exc) {
			if exc.ErrorCode() == "WAFNonexistentItemException" {
				return nil
			}
		}

		return err
	}
	if response.WebACL == nil {
		return nil
	}

	return resource.Set(c.Name, response.WebACL.ARN)
}

func getLoadBalancer(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	lb := resource.Item.(*ohaws.LoadBalancerV2)

	attrsResp, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: lb.LoadBalancerArn})
	if err != nil {
		return fmt.Errorf("error describing load balancer attributes: %w", err)
	}

	attrsMap := make(map[string]string, len(attrsResp.Attributes))
	for i := range attrsResp.Attributes {
		if s := aws.ToString(attrsResp.Attributes[i].Key); s != "" {
			attrsMap[s] = aws.ToString(attrsResp.Attributes[i].Value)
		}
	}

	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &lb})
	if err != nil {
		return fmt.Errorf("error instantiating decoder: %w", err)
	}
	if err := d.Decode(attrsMap); err != nil {
		return fmt.Errorf("error decoding load balancer attributes: %w", err)
	}

	tagsResp, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{ResourceArns: []string{aws.ToString(lb.LoadBalancerArn)}})
	if err != nil {
		return fmt.Errorf("error describing load balancer tags: %w", err)
	}

	arnStr := aws.ToString(lb.LoadBalancerArn)
	for _, t := range tagsResp.TagDescriptions {
		if aws.ToString(t.ResourceArn) == arnStr {
			lb.Tags = client.TagsToMap(t.Tags)
		}
	}

	resource.Item = lb

	return nil
}
