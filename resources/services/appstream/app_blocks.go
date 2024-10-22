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

func AppBlocks() *schema.Table {
	tableName := "aws_appstream_app_blocks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html`,
		Resolver:    fetchAppstreamAppBlocks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   transformers.TransformWithStruct(&types.AppBlock{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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

func fetchAppstreamAppBlocks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeAppBlocksInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppstream).Appstream
	// No paginator available
	for {
		response, err := svc.DescribeAppBlocks(ctx, &input, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.AppBlocks

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
