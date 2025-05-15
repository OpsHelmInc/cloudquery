package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElbv2Listeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	lb := parent.Item.(*ohaws.LoadBalancerV2)
	config := elbv2.DescribeListenersInput{
		LoadBalancerArn: lb.LoadBalancerArn,
	}
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeListeners(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		listeners := make([]*ohaws.LoadBalancerV2Listener, len(response.Listeners))
		for idx := range response.Listeners {
			listeners[idx] = &ohaws.LoadBalancerV2Listener{
				Listener: response.Listeners[idx],
			}
		}
		res <- listeners
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func getLoadBalancerListener(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticloadbalancingv2
	listener := resource.Item.(*ohaws.LoadBalancerV2Listener)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*listener.ListenerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	listener.Tags = make(map[string]string, len(tagsOutput.TagDescriptions))
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			listener.Tags[aws.ToString(s.Key)] = aws.ToString(s.Value)
		}
	}

	return nil
}

func fetchElbv2ListenerCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	region := c.Region
	svc := c.Services().Elasticloadbalancingv2
	listener := parent.Item.(*ohaws.LoadBalancerV2Listener)
	config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
	for {
		response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
			options.Region = region
		})
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
