package fsx

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFsxSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Fsx
	input := fsx.DescribeSnapshotsInput{MaxResults: aws.Int32(1000)}
	paginator := fsx.NewDescribeSnapshotsPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- result.Snapshots
	}
	return nil
}
