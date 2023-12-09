package kafka

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
)

func fetchKafkaClusterOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	if parent.Item.(*types.Cluster).ClusterType == types.ClusterTypeServerless {
		// serverless clusters do not support cluster operations
		return nil
	}

	var input = getListClusterOperationsInput(parent)
	c := meta.(*client.Client)
	svc := c.Services().Kafka
	for {
		response, err := svc.ListClusterOperations(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.ClusterOperationInfoList
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
