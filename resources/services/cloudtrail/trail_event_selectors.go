package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/cloudtrail/models"
)

func trailEventSelectors() *schema.Table {
	tableName := "aws_cloudtrail_trail_event_selectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_GetEventSelectors.html`,
		Resolver:    fetchCloudtrailTrailEventSelectors,
		Transform: transformers.TransformWithStruct(&cloudtrail.GetEventSelectorsOutput{},
			transformers.WithPrimaryKeyComponents("TrailARN"),
			transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false)},
	}
}

func fetchCloudtrailTrailEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.CloudTrailWrapper)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudtrail).Cloudtrail
	response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
		options.Region = *r.HomeRegion
	})
	if err != nil {
		return err
	}
	res <- response
	return nil
}
