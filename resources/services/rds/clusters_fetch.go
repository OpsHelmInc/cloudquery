package rds

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func fetchRdsClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeDBClustersInput
	c := meta.(*client.Client)
	svc := c.Services().Rds
	for {
		response, err := svc.DescribeDBClusters(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DBClusters
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
func resolveRdsClusterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.DBCluster)
	tags := map[string]*string{}
	for _, t := range r.TagList {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
