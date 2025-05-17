package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElasticacheReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(meta.(*client.Client).Services().Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		groups := make([]*ohaws.ReplicationGroup, len(v.ReplicationGroups))
		for idx, g := range v.ReplicationGroups {
			groups[idx] = &ohaws.ReplicationGroup{
				ReplicationGroup: g,
			}
		}
		res <- groups
	}
	return nil
}

func getReplicationGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	group := resource.Item.(*ohaws.ReplicationGroup)
	svc := c.Services().Elasticache

	tagsResp, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: group.ARN,
	})
	if err != nil {
		return err
	}

	group.Tags = make([]ohaws.Tag, len(tagsResp.TagList))
	for i := range tagsResp.TagList {
		group.Tags[i].Key = tagsResp.TagList[i].Key
		group.Tags[i].Value = tagsResp.TagList[i].Value
	}

	return nil
}
