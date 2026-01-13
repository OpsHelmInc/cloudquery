package route53

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRoute53TrafficPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListTrafficPoliciesInput
	c := meta.(*client.Client)
	svc := c.Services().Route53

	for {
		if err := c.WaitForRateLimit(ctx, serviceName); err != nil {
			return err
		}
		response, err := svc.ListTrafficPolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrafficPolicySummaries

		if aws.ToString(response.TrafficPolicyIdMarker) == "" {
			break
		}
		config.TrafficPolicyIdMarker = response.TrafficPolicyIdMarker
	}
	return nil
}
func fetchRoute53TrafficPolicyVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.TrafficPolicySummary)
	c := meta.(*client.Client)
	config := route53.ListTrafficPolicyVersionsInput{Id: r.Id}
	svc := c.Services().Route53
	for {
		if err := c.WaitForRateLimit(ctx, serviceName); err != nil {
			return err
		}
		response, err := svc.ListTrafficPolicyVersions(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.TrafficPolicies
		if aws.ToString(response.TrafficPolicyVersionMarker) == "" {
			break
		}
		config.TrafficPolicyVersionMarker = response.TrafficPolicyVersionMarker
	}
	return nil
}
func resolveTrafficPolicyArn() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
		return []string{"trafficpolicy", *resource.Item.(types.TrafficPolicySummary).Id}, nil
	})
}
