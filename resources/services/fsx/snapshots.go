package fsx

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Snapshots() *schema.Table {
	tableName := "aws_fsx_snapshots"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchFsxSnapshots,
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ResourceARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "administrative_actions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("AdministrativeActions"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchFsxSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceFsx).Fsx
	input := fsx.DescribeSnapshotsInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *fsx.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- result.Snapshots
	}
	return nil
}
