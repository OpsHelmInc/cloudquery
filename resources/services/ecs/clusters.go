package ecs

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Clusters() *schema.Table {
	tableName := "aws_ecs_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchEcsClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ClusterArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			clusterTasks(),
			clusterServices(),
			clusterContainerInstances(),
		},
	}
}

func fetchEcsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListClustersInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcs).Ecs
	paginator := ecs.NewListClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.ClusterArns) == 0 {
			return nil
		}
		describeClusterOutput, err := svc.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: page.ClusterArns,
			Include: []types.ClusterField{
				types.ClusterFieldAttachments,
				types.ClusterFieldTags,
				types.ClusterFieldSettings,
				types.ClusterFieldConfigurations,
				types.ClusterFieldStatistics,
			},
		}, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- describeClusterOutput.Clusters
	}
	return nil
}
