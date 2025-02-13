package stepfunctions

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Activities() *schema.Table {
	tableName := "aws_stepfunctions_activities"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/step-functions/latest/apireference/API_ListActivities.html`,
		Resolver:    fetchStepfunctionsActivities,
		Transform:   transformers.TransformWithStruct(&types.ActivityListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ActivityArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchStepfunctionsActivities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSfn).Sfn
	config := sfn.ListActivitiesInput{
		MaxResults: 1000,
	}
	paginator := sfn.NewListActivitiesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *sfn.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Activities
	}
	return nil
}
