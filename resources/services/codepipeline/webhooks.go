package codepipeline

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Webhooks() *schema.Table {
	tableName := "aws_codepipeline_webhooks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html`,
		Resolver:    fetchCodepipelineWebhooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codepipeline"),
		Transform:   transformers.TransformWithStruct(&types.ListWebhookItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
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

func fetchCodepipelineWebhooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodepipeline).Codepipeline
	paginator := codepipeline.NewListWebhooksPaginator(svc, &codepipeline.ListWebhooksInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codepipeline.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Webhooks
	}
	return nil
}
