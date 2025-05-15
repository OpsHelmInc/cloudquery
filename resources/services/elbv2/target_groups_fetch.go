package elbv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElbv2TargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elbv2.DescribeTargetGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Elasticloadbalancingv2
	for {
		response, err := svc.DescribeTargetGroups(ctx, &config)
		if err != nil {
			return err
		}
		tgs := make([]*ohaws.TargetGroup, len(response.TargetGroups))
		for idx := range response.TargetGroups {
			tgs[idx] = &ohaws.TargetGroup{
				TargetGroup: response.TargetGroups[idx],
			}
		}
		res <- tgs
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}

func getTargetGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	region := meta.(*client.Client).Region
	svc := meta.(*client.Client).Services().Elasticloadbalancingv2
	targetGroup := resource.Item.(*ohaws.TargetGroup)

	attrsResp, err := svc.DescribeTargetGroupAttributes(ctx, &elbv2.DescribeTargetGroupAttributesInput{
		TargetGroupArn: targetGroup.TargetGroupArn,
	})
	if err != nil {
		return fmt.Errorf("error target group attributes: %w", err)
	}

	attrsMap := make(map[string]string, len(attrsResp.Attributes))
	for i := range attrsResp.Attributes {
		if s := aws.ToString(attrsResp.Attributes[i].Key); s != "" {
			attrsMap[s] = aws.ToString(attrsResp.Attributes[i].Value)
		}
	}

	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &targetGroup})
	if err != nil {
		return fmt.Errorf("error instantiating decoder: %w", err)
	}
	if err := d.Decode(attrsMap); err != nil {
		return fmt.Errorf("error decoding target group attributes: %w", err)
	}

	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*targetGroup.TargetGroupArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}

	targetGroup.Tags = make(map[string]string, len(tagsOutput.TagDescriptions))
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			targetGroup.Tags[aws.ToString(s.Key)] = aws.ToString(s.Value)
		}
	}

	return nil
}

func fetchElbv2TargetGroupTargetHealthDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Elasticloadbalancingv2
	tg := parent.Item.(*ohaws.TargetGroup)
	response, err := svc.DescribeTargetHealth(ctx, &elbv2.DescribeTargetHealthInput{
		TargetGroupArn: tg.TargetGroupArn,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.TargetHealthDescriptions
	return nil
}
