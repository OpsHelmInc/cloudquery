package cloudfront

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func CachePolicies() *schema.Table {
	tableName := "aws_cloudfront_cache_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CachePolicySummary.html`,
		Resolver:    fetchCloudfrontCachePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:   transformers.TransformWithStruct(&types.CachePolicySummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("CachePolicy.Id"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveCachePolicyARN(),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListCachePoliciesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudfront).Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, nil, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if response.CachePolicyList != nil {
			res <- response.CachePolicyList.Items
		}

		if aws.ToString(response.CachePolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.CachePolicyList.NextMarker
	}
	return nil
}

func resolveCachePolicyARN() schema.ColumnResolver {
	return client.ResolveARNWithAccount(client.CloudfrontService, func(resource *schema.Resource) ([]string, error) {
		return []string{"cache-policy", *resource.Item.(types.CachePolicySummary).CachePolicy.Id}, nil
	})
}
