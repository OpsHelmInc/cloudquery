package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/thoas/go-funk"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func getListNodesInput(parent *schema.Resource) kafka.ListNodesInput {
	return kafka.ListNodesInput{
		ClusterArn: parent.Item.(*types.Cluster).ClusterArn,
		MaxResults: aws.Int32(100),
	}
}

func getListClusterOperationsInput(parent *schema.Resource) kafka.ListClusterOperationsInput {
	return kafka.ListClusterOperationsInput{
		ClusterArn: parent.Item.(*types.Cluster).ClusterArn,
		MaxResults: aws.Int32(100),
	}
}

func describeClustersInput(resource *schema.Resource) kafka.DescribeClusterV2Input {
	return kafka.DescribeClusterV2Input{
		ClusterArn: resource.Item.(types.Cluster).ClusterArn,
	}
}

func resolveKafkaTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		if path == "OperationArn" {
			// cluster operations do not support tags. In a future release we should remove the column from the `aws_kafka_cluster_operations` table
			return r.Set(c.Name, map[string]string{})
		}
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services(client.AWSServiceKafka).Kafka
		params := kafka.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params, func(options *kafka.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, output.Tags)
	}
}
