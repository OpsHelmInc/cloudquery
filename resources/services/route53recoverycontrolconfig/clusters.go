package route53recoverycontrolconfig

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Clusters() *schema.Table {
	tableName := "aws_route53recoverycontrolconfig_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/recovery-cluster/latest/api/cluster.html`,
		Resolver:    fetchClusters,
		Transform:   transformers.TransformWithStruct(&types.Cluster{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53-recovery-control-config"),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ClusterArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoverycontrolconfig).Route53recoverycontrolconfig
	paginator := route53recoverycontrolconfig.NewListClustersPaginator(svc, &route53recoverycontrolconfig.ListClustersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoverycontrolconfig.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Clusters
	}
	return nil
}
