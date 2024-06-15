package elbv1

import (
	"context"

	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElbv1LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancing
	processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
		tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
		for _, lb := range loadBalancers {
			tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
		}
		tagsResponse, err := svc.DescribeTags(ctx, tagsCfg, func(options *elbv1.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, lb := range loadBalancers {
			loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName}, func(options *elbv1.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}

			wrapper := ohaws.LoadBalancerV1{
				LoadBalancerDescription: lb,
				Tags:                    client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
				LoadBalancerAttributes:  loadBalancerAttributes.LoadBalancerAttributes,
			}

			res <- wrapper
		}
		return nil
	}
	paginator := elbv1.NewDescribeLoadBalancersPaginator(svc, &elbv1.DescribeLoadBalancersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *elbv1.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		for i := 0; i < len(page.LoadBalancerDescriptions); i += 20 {
			end := i + 20

			if end > len(page.LoadBalancerDescriptions) {
				end = len(page.LoadBalancerDescriptions)
			}
			loadBalancers := page.LoadBalancerDescriptions[i:end]
			if err := processLoadBalancers(loadBalancers); err != nil {
				return err
			}
		}
	}

	return nil
}

func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(ohaws.LoadBalancerV1)
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancing
	response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName})
	if err != nil {
		return err
	}
	res <- response.PolicyDescriptions
	return nil
}

func resolveElbv1loadBalancerPolicyAttributeDescriptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyDescription)

	response := make(map[string]any, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}

func resolveLoadBalancerARN() schema.ColumnResolver {
	return client.ResolveARN(client.ElasticLoadBalancingService, func(resource *schema.Resource) ([]string, error) {
		return []string{"loadbalancer", *resource.Item.(ohaws.LoadBalancerV1).LoadBalancerName}, nil
	})
}
