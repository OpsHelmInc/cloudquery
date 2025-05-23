package wellarchitected

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func lensReviewImprovements() *schema.Table {
	name := "aws_wellarchitected_lens_review_improvements"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_ImprovementSummary.html`,
		Transform: transformers.TransformWithStruct(new(types.ImprovementSummary),
			transformers.WithPrimaryKeyComponents("PillarId", "QuestionId"),
		),
		Resolver: fetchLensReviewImprovements,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "workload_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("workload_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "workload_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("workload_id"),
			},
			{
				Name:                "milestone_number",
				Type:                arrow.PrimitiveTypes.Int64,
				Resolver:            schema.ParentColumnResolver("milestone_number"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "lens_alias",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("lens_alias"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLensReviewImprovements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	review, ok := parent.Item.(*types.LensReview)
	if !ok {
		// we need the full resource to get the pillars
		return nil
	}

	cl := meta.(*client.Client)
	service := cl.Services(client.AWSServiceWellarchitected).Wellarchitected
	milestoneNumber := int32(parent.Get("milestone_number").Get().(int64))
	workloadID := parent.Get("workload_id").String()

	for _, pillar := range review.PillarReviewSummaries {
		p := wellarchitected.NewListLensReviewImprovementsPaginator(service,
			&wellarchitected.ListLensReviewImprovementsInput{
				LensAlias:       review.LensAlias,
				WorkloadId:      &workloadID,
				MilestoneNumber: aws.Int32(milestoneNumber),
				MaxResults:      aws.Int32(50),
				PillarId:        pillar.PillarId,
			},
		)
		for p.HasMorePages() {
			output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}

			res <- output.ImprovementSummaries
		}
	}

	return nil
}
