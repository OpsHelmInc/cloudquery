package cloudformation

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	cqtypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func templateSummaries() *schema.Table {
	tableName := "aws_cloudformation_template_summaries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplateSummary.html`,
		Resolver:    fetchTemplateSummary,
		Transform: transformers.TransformWithStruct(
			&cloudformation.GetTemplateSummaryOutput{},
			transformers.WithSkipFields("ResultMetadata"), // This field contains metadata about the API call rather than the template itself, so remove it.
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "stack_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:                "stack_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "metadata",
				Type:     cqtypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Metadata"),
			},
		},
	}
}

func fetchTemplateSummary(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudformation).Cloudformation

	stack := parent.Item.(types.Stack)

	summary, err := svc.GetTemplateSummary(ctx, &cloudformation.GetTemplateSummaryInput{
		StackName: stack.StackName,
	}, func(o *cloudformation.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- summary
	return nil
}
