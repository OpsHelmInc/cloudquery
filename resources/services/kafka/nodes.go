package kafka

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func nodes() *schema.Table {
	tableName := "aws_kafka_nodes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-nodes.html#ListNodes`,
		Resolver:    fetchKafkaNodes,
		Transform:   transformers.TransformWithStruct(&types.NodeInfo{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("NodeARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "cluster_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchKafkaNodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	if parent.Item.(*types.Cluster).ClusterType == types.ClusterTypeServerless {
		// serverless clusters do not have nodes
		return nil
	}
	if parent.Item.(*types.Cluster).State == types.ClusterStateCreating {
		// nodes for clusters in state "CREATING" cannot be listed
		return nil
	}

	var input = getListNodesInput(parent)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceKafka).Kafka
	paginator := kafka.NewListNodesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kafka.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.NodeInfoList
	}
	return nil
}
