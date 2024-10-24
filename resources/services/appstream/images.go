package appstream

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Images() *schema.Table {
	tableName := "aws_appstream_images"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Image.html`,
		Resolver:    fetchAppstreamImages,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.Image{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchAppstreamImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := appstream.DescribeImagesInput{MaxResults: aws.Int32(25)}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppstream).Appstream
	paginator := appstream.NewDescribeImagesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Images
	}
	return nil
}
