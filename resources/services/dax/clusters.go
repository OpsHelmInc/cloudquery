package dax

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Clusters() *schema.Table {
	tableName := "aws_dax_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html`,
		Resolver:    fetchDaxClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dax"),
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
				Resolver: resolveClusterTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchDaxClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDax).Dax
	config := dax.DescribeClustersInput{}
	// No paginator available
	for {
		output, err := svc.DescribeClusters(ctx, &config, func(options *dax.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- output.Clusters

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}

func resolveClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cluster := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDax).Dax
	input := &dax.ListTagsInput{
		ResourceName: cluster.ClusterArn,
	}
	var tags []types.Tag
	// No paginator available
	for {
		response, err := svc.ListTags(ctx, input, func(options *dax.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, response.Tags...)
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
