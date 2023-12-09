package lightsail

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/resources/services/lightsail/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"golang.org/x/sync/errgroup"
)

func fetchLightsailDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetDistributions(ctx, &input, func(options *lightsail.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
		})
		if err != nil {
			return err
		}

		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(MaxGoroutines)
		for _, d := range response.Distributions {
			func(d types.LightsailDistribution) {
				errs.Go(func() error {
					return fetchCacheReset(ctx, res, c, d)
				})
			}(d)
		}
		err = errs.Wait()
		if err != nil {
			return err
		}
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}

func fetchCacheReset(ctx context.Context, res chan<- interface{}, c *client.Client, d types.LightsailDistribution) error {
	svc := c.Services().Lightsail
	resetInput := lightsail.GetDistributionLatestCacheResetInput{
		DistributionName: d.Name,
	}
	resetResp, err := svc.GetDistributionLatestCacheReset(ctx, &resetInput, func(options *lightsail.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}
	res <- models.DistributionWrapper{LightsailDistribution: &d, LatestCacheReset: resetResp}
	return nil
}
