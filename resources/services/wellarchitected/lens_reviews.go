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

func lensReviews() *schema.Table {
	name := "aws_wellarchitected_lens_reviews"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_LensReview.html`,
		Transform: transformers.TransformWithStruct(new(types.LensReview),
			transformers.WithPrimaryKeyComponents("LensAlias"),
		),
		Resolver: fetchLensReviews,
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
		},
		Relations: schema.Tables{lensReviewImprovements()},
	}
}

func fetchLensReviews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services(client.AWSServiceWellarchitected).Wellarchitected
	milestoneNumber := int32(parent.Get("milestone_number").Get().(int64))
	workloadID := parent.Get("workload_id").String()

	p := wellarchitected.NewListLensReviewsPaginator(service,
		&wellarchitected.ListLensReviewsInput{
			WorkloadId:      &workloadID,
			MilestoneNumber: aws.Int32(milestoneNumber),
			MaxResults:      aws.Int32(50),
		},
	)
	for p.HasMorePages() {
		output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}

		// we also need workload summary for Get call, so we do it here
		for _, summary := range output.LensReviewSummaries {
			out, err := service.GetLensReview(ctx,
				&wellarchitected.GetLensReviewInput{
					LensAlias:       summary.LensAlias,
					WorkloadId:      &workloadID,
					MilestoneNumber: aws.Int32(milestoneNumber),
				},
				func(o *wellarchitected.Options) { o.Region = cl.Region },
			)
			if err != nil {
				cl.Logger().Err(err).Str("table", "aws_wellarchitected_lens_reviews").Msg("Failed to perform get, ignoring...")
				// At the very least we want the summary data to be filled in
				res <- summary
				continue
			}

			res <- out.LensReview
		}
	}

	return nil
}
