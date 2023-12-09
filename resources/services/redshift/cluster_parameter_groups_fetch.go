package redshift

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

func fetchRedshiftClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cluster := parent.Item.(types.Cluster)
	res <- cluster.ClusterParameterGroups
	return nil
}
