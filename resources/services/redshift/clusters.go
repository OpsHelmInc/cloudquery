package redshift

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Clusters() *schema.Table {
	tableName := "aws_redshift_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchClusters,
		Transform:   transformers.TransformWithStruct(&types.Cluster{}, transformers.WithSkipFields("ClusterParameterGroups")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveClusterArn(),
				Description:         `The Amazon Resource Name (ARN) for the resource.`,
				PrimaryKeyComponent: true,
			},
			{
				Name:        "logging_status",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveRedshiftClusterLoggingStatus,
				Description: `Describes the status of logging for a cluster.`,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			snapshots(),
			clusterParameterGroups(),
			endpointAccess(),
			endpointAuthorization(),
		},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config redshift.DescribeClustersInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRedshift).Redshift
	paginator := redshift.NewDescribeClustersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Clusters
	}
	return nil
}

func resolveRedshiftClusterLoggingStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Cluster)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRedshift).Redshift
	cfg := redshift.DescribeLoggingStatusInput{
		ClusterIdentifier: r.ClusterIdentifier,
	}
	response, err := svc.DescribeLoggingStatus(ctx, &cfg, func(options *redshift.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	return resource.Set(c.Name, response)
}

func resolveClusterArn() schema.ColumnResolver {
	return client.ResolveARN(client.RedshiftService, func(resource *schema.Resource) ([]string, error) {
		return []string{fmt.Sprintf("cluster:%s", *resource.Item.(types.Cluster).ClusterIdentifier)}, nil
	})
}
