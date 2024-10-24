package elasticache

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ReplicationGroups() *schema.Table {
	tableName := "aws_elasticache_replication_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ReplicationGroup.html`,
		Resolver:    fetchElasticacheReplicationGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.ReplicationGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveElasticacheReplicationGroupTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchElasticacheReplicationGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := elasticache.NewDescribeReplicationGroupsPaginator(meta.(*client.Client).Services(client.AWSServiceElasticache).Elasticache, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.ReplicationGroups
	}
	return nil
}

func resolveElasticacheReplicationGroupTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticache).Elasticache
	tags, err := svc.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{ResourceName: r.Item.(types.ReplicationGroup).ARN}, func(options *elasticache.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return r.Set(c.Name, client.TagsToMap(tags.TagList))
}
