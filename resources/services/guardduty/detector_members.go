package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/guardduty/models"
)

func detectorMembers() *schema.Table {
	tableName := "aws_guardduty_detector_members"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/guardduty/latest/APIReference/API_Member.html`,
		Resolver:    fetchDetectorMembers,
		Transform: transformers.TransformWithStruct(&types.Member{},
			transformers.WithPrimaryKeyComponents("AccountId"),
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

func fetchDetectorMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	detector := parent.Item.(models.DetectorWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGuardduty).Guardduty
	config := &guardduty.ListMembersInput{DetectorId: aws.String(detector.Id)}
	paginator := guardduty.NewListMembersPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *guardduty.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Members
	}
	return nil
}
