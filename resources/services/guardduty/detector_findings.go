package guardduty

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/resources/services/guardduty/models"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func detectorFindings() *schema.Table {
	tableName := "aws_guardduty_detector_findings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Finding.html`,
		Resolver:    fetchDetectorFindings,
		Transform: transformers.TransformWithStruct(&types.Finding{},
			transformers.WithPrimaryKeyComponents("Arn"),
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
		),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			detectorARNColumn,
		},
	}
}

func fetchDetectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	config := &guardduty.ListFindingsInput{
		DetectorId: &detector.Id,
	}
	paginator := guardduty.NewListFindingsPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.FindingIds) == 0 {
			continue
		}

		f, err := svc.GetFindings(ctx, &guardduty.GetFindingsInput{
			DetectorId: &detector.Id,
			FindingIds: page.FindingIds,
		}, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- f.Findings
	}
	return nil
}
