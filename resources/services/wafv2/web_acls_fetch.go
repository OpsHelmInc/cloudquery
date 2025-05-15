package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	service := c.Services().Wafv2

	config := wafv2.ListWebACLsInput{
		Scope: c.WAFScope,
		Limit: aws.Int32(100),
	}
	for {
		output, err := service.ListWebACLs(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.WebACLs

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}

func getWebAcl(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Wafv2
	webAcl := resource.Item.(types.WebACLSummary)

	webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: c.WAFScope}
	webAclOutput, err := svc.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}

	cfg := wafv2.GetLoggingConfigurationInput{
		ResourceArn: webAclOutput.WebACL.ARN,
	}

	loggingConfigurationOutput, err := svc.GetLoggingConfiguration(ctx, &cfg, func(options *wafv2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		if client.IsAWSError(err, "WAFNonexistentItemException") {
			c.Logger().Debug().Err(err).Msg("Logging configuration not found for")
		} else {
			c.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed with error")
		}
	}

	tags, err := getWebACLTags(ctx, c, svc, *webAcl.ARN)
	if err != nil {
		return err
	}

	var webAclLoggingConfiguration *types.LoggingConfiguration
	if loggingConfigurationOutput != nil {
		webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
	}

	resource.Item = &ohaws.WAFv2WebACL{
		WebACL:               *webAclOutput.WebACL,
		LoggingConfiguration: webAclLoggingConfiguration,
		Tags:                 tags,
	}
	return nil
}

func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*ohaws.WAFv2WebACL)

	cl := meta.(*client.Client)
	service := cl.Services().Wafv2

	resourceArns := []string{}
	if cl.WAFScope == types.ScopeCloudfront {
		cloudfrontService := cl.Services().Cloudfront
		params := &cloudfront.ListDistributionsByWebACLIdInput{
			WebACLId: webACL.Id,
			MaxItems: aws.Int32(100),
		}
		for {
			output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(options *cloudfront.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for _, item := range output.DistributionList.Items {
				resourceArns = append(resourceArns, *item.ARN)
			}
			if aws.ToString(output.DistributionList.NextMarker) == "" {
				break
			}
			params.Marker = output.DistributionList.NextMarker
		}
	} else {
		output, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN})
		if err != nil {
			return err
		}
		resourceArns = output.ResourceArns
	}
	return resource.Set(c.Name, resourceArns)
}

func getWebACLTags(ctx context.Context, c *client.Client, svc services.Wafv2Client, webACLArn string) ([]types.Tag, error) {
	var tags []types.Tag
	params := wafv2.ListTagsForResourceInput{ResourceARN: aws.String(webACLArn)}
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
