package batch

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func JobDefinitions() *schema.Table {
	tableName := "aws_batch_job_definitions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/batch/latest/APIReference/API_DescribeJobDefinitions.html`,
		Resolver:    fetchBatchJobDefinitions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "batch"),
		Transform:   transformers.TransformWithStruct(&types.JobDefinition{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveBatchJobDefinitionTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("JobDefinitionArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchBatchJobDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := batch.DescribeJobDefinitionsInput{
		MaxResults: aws.Int32(100),
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBatch).Batch
	p := batch.NewDescribeJobDefinitionsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *batch.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.JobDefinitions
	}
	return nil
}

func resolveBatchJobDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBatch).Batch
	summary := resource.Item.(types.JobDefinition)

	input := batch.ListTagsForResourceInput{
		ResourceArn: summary.JobDefinitionArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input, func(options *batch.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.Tags)
}
