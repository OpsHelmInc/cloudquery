package wafv2

import (
	"context"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchWafv2Ipsets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2

	params := wafv2.ListIPSetsInput{
		Scope: cl.WAFScope,
		Limit: aws.Int32(100), // maximum value: https://docs.aws.amazon.com/waf/latest/APIReference/API_ListIPSets.html
	}
	for {
		result, err := svc.ListIPSets(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.IPSets

		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func getIpset(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafv2
	s := resource.Item.(types.IPSetSummary)

	info, err := svc.GetIPSet(
		ctx,
		&wafv2.GetIPSetInput{
			Id:    s.Id,
			Name:  s.Name,
			Scope: cl.WAFScope,
		},
		func(options *wafv2.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	tags, err := getIpsetTags(ctx, cl, svc, *info.IPSet.ARN)
	if err != nil {
		return err
	}

	resource.Item = &ohaws.WAFv2IPSet{
		IPSet: *info.IPSet,
		Tags:  tags,
	}

	return nil
}

func resolveIpsetAddresses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(*ohaws.WAFv2IPSet)
	addrs := make([]*net.IPNet, 0, len(s.Addresses))
	for _, a := range s.Addresses {
		_, n, err := net.ParseCIDR(a)
		if err != nil {
			return err
		}
		addrs = append(addrs, n)
	}
	return resource.Set(c.Name, addrs)
}

func getIpsetTags(ctx context.Context, c *client.Client, svc services.Wafv2Client, ipsetArn string) ([]types.Tag, error) {
	var tags []types.Tag
	params := wafv2.ListTagsForResourceInput{ResourceARN: aws.String(ipsetArn)}
	for {
		result, err := svc.ListTagsForResource(ctx, &params, func(options *wafv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return nil, err
		}
		if result != nil || result.TagInfoForResource != nil {
			tags = append(tags, result.TagInfoForResource.TagList...)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}

	return tags, nil
}
