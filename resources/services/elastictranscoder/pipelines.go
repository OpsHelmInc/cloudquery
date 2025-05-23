package elastictranscoder

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Pipelines() *schema.Table {
	tableName := "aws_elastictranscoder_pipelines"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-pipelines.html`,
		Resolver:    fetchElastictranscoderPipelines,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elastictranscoder"),
		Transform:   transformers.TransformWithStruct(&types.Pipeline{}),
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

		Relations: []*schema.Table{
			pipelineJobs(),
		},
	}
}

func fetchElastictranscoderPipelines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElastictranscoder).Elastictranscoder

	p := elastictranscoder.NewListPipelinesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *elastictranscoder.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- response.Pipelines
	}

	return nil
}
