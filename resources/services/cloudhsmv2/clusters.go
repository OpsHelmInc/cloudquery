package cloudhsmv2

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Clusters() *schema.Table {
	tableName := "aws_cloudhsmv2_clusters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchCloudhsmv2Clusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudhsmv2"),
		Transform:   transformers.TransformWithStruct(&types.Cluster{}, transformers.WithSkipFields("TagList")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveClusterArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagPath("TagList"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchCloudhsmv2Clusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudhsmv2).Cloudhsmv2
	var input cloudhsmv2.DescribeClustersInput
	paginator := cloudhsmv2.NewDescribeClustersPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *cloudhsmv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Clusters
	}
	return nil
}

func resolveClusterArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Cluster)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "hsm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "cluster/" + aws.ToString(item.ClusterId),
	}
	return resource.Set(c.Name, a.String())
}
