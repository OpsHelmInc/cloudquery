package rds

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func clusterBacktracks() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_backtracks",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBClusterBacktracks.html`,
		Resolver:    fetchRdsClusterBacktracks,
		Transform: transformers.TransformWithStruct(
			&types.DBClusterBacktrack{},
			transformers.WithPrimaryKeyComponents("BacktrackIdentifier"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "db_cluster_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchRdsClusterBacktracks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(types.DBCluster)

	if aws.ToInt64(cluster.BacktrackWindow) == 0 {
		// If this value is set to 0, backtracking is disabled for the DB cluster.
		return nil
	}

	config := rds.DescribeDBClusterBacktracksInput{
		DBClusterIdentifier: cluster.DBClusterIdentifier,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	p := rds.NewDescribeDBClusterBacktracksPaginator(svc, &config)
	for p.HasMorePages() {
		resp, err := p.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- resp.DBClusterBacktracks
	}
	return nil
}
