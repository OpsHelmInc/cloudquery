package rds

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ClusterSnapshots() *schema.Table {
	tableName := "aws_rds_cluster_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html`,
		Resolver:    fetchRdsClusterSnapshots,
		Transform:   transformers.TransformWithStruct(&types.DBClusterSnapshot{}, transformers.WithSkipFields("TagList")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBClusterSnapshotArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTagPath("TagList"),
			},
			{
				Name:     "attributes",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRDSClusterSnapshotAttributes,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchRdsClusterSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	var input rds.DescribeDBClusterSnapshotsInput
	paginator := rds.NewDescribeDBClusterSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return nil
		}
		res <- page.DBClusterSnapshots
	}
	return nil
}

func resolveRDSClusterSnapshotAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	s := resource.Item.(types.DBClusterSnapshot)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	out, err := svc.DescribeDBClusterSnapshotAttributes(
		ctx,
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		func(o *rds.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if out.DBClusterSnapshotAttributesResult == nil {
		return nil
	}

	return resource.Set(column.Name, out.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes)
}
