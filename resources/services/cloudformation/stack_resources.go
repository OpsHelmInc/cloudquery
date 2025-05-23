package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func stackResources() *schema.Table {
	tableName := "aws_cloudformation_stack_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_StackResourceSummary.html`,
		Resolver:    fetchCloudformationStackResources,
		Transform:   transformers.TransformWithStruct(&types.StackResourceSummary{}, transformers.WithPrimaryKeyComponents("LogicalResourceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "stack_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{StackName: stack.StackId} // the docs suggest that stack ID is better
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudformation).Cloudformation
	paginator := cloudformation.NewListStackResourcesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudformation.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.StackResourceSummaries
	}
	return nil
}
