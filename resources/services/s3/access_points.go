package s3

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func AccessPoints() *schema.Table {
	tableName := "aws_s3_access_points"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AccessPoint.html",
		Resolver:    fetchAccessPoints,
		Transform:   transformers.TransformWithStruct(&types.AccessPoint{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "s3-control"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("AccessPointArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchAccessPoints(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3control).S3control

	paginator := s3control.NewListAccessPointsPaginator(svc, &s3control.ListAccessPointsInput{
		AccountId: aws.String(cl.AccountID),
	})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *s3control.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.AccessPointList
	}

	return nil
}
