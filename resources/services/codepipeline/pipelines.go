package codepipeline

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Pipelines() *schema.Table {
	tableName := "aws_codepipeline_pipelines"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html`,
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codepipeline"),
		Transform:           transformers.TransformWithStruct(&codepipeline.GetPipelineOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolvePipelineArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolvePipelineTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchCodepipelinePipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodepipeline).Codepipeline
	config := codepipeline.ListPipelinesInput{}
	paginator := codepipeline.NewListPipelinesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codepipeline.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Pipelines
	}
	return nil
}

func getPipeline(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodepipeline).Codepipeline
	item := resource.Item.(types.PipelineSummary)
	response, err := svc.GetPipeline(ctx, &codepipeline.GetPipelineInput{Name: item.Name}, func(options *codepipeline.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response
	return nil
}

func resolvePipelineTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodepipeline).Codepipeline
	paginator := codepipeline.NewListTagsForResourcePaginator(svc, &codepipeline.ListTagsForResourceInput{
		ResourceArn: pipeline.Metadata.PipelineArn,
	})
	var tags []types.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codepipeline.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}

func resolvePipelineArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	pipeline := resource.Item.(*codepipeline.GetPipelineOutput)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "codepipeline",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "pipelines/" + *pipeline.Metadata.PipelineArn,
	}

	return resource.Set(c.Name, a.String())
}
