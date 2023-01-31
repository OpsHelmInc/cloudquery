package elasticache

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticacheParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	awsProviderClient := meta.(*client.Client)
	svc := awsProviderClient.Services().Elasticache

	var describeCacheParameterGroupsInput elasticache.DescribeCacheParameterGroupsInput

	for {
		describeCacheParameterGroupsOutput, err := svc.DescribeCacheParameterGroups(ctx, &describeCacheParameterGroupsInput)

		if err != nil {
			return err
		}

		res <- describeCacheParameterGroupsOutput.CacheParameterGroups

		if aws.ToString(describeCacheParameterGroupsOutput.Marker) == "" {
			return nil
		}

		describeCacheParameterGroupsInput.Marker = describeCacheParameterGroupsOutput.Marker
	}
}
