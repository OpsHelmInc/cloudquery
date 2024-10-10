package emr

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func notebookExecutions() *schema.Table {
	tableName := "aws_emr_notebook_executions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_NotebookExecution.html`,
		Resolver:            fetchNotebookExecutions,
		PreResourceResolver: getNotebookExecution,
		Transform:           transformers.TransformWithStruct(&types.NotebookExecution{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "cluster_arn",
				Description:         "The Amazon Resource Name (ARN) of the EMR Cluster.",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: false,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchNotebookExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	p := parent.Item.(*types.Cluster)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListNotebookExecutionsPaginator(svc, &emr.ListNotebookExecutionsInput{ExecutionEngineId: p.Id})
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.NotebookExecutions
	}
	return nil
}

func getNotebookExecution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	response, err := svc.DescribeNotebookExecution(ctx, &emr.DescribeNotebookExecutionInput{NotebookExecutionId: resource.Item.(types.NotebookExecutionSummary).NotebookExecutionId}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.NotebookExecution
	err = resource.Set("tags", client.TagsToMap(response.NotebookExecution.Tags))
	if err != nil {
		return err
	}
	return nil
}
