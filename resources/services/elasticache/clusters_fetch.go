package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElasticacheClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeCacheClustersInput
	input.ShowCacheNodeInfo = aws.Bool(true)

	paginator := elasticache.NewDescribeCacheClustersPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		clusters := make([]*ohaws.CacheCluster, len(v.CacheClusters))
		for idx, c := range v.CacheClusters {
			clusters[idx] = &ohaws.CacheCluster{
				CacheCluster: c,
			}
		}
		res <- clusters
	}
	return nil
}

func getCluster(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	cluster := resource.Item.(*ohaws.CacheCluster)
	svc := c.Services().Elasticache

	tagsResp, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: cluster.ARN,
	})
	if err != nil {
		return err
	}

	cluster.Tags = make([]ohaws.Tag, len(tagsResp.TagList))
	for i := range tagsResp.TagList {
		cluster.Tags[i].Key = tagsResp.TagList[i].Key
		cluster.Tags[i].Value = tagsResp.TagList[i].Value
	}

	return nil
}
