package eventbridge

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func ApiDestinations() *schema.Table {
	tableName := "aws_eventbridge_api_destinations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html`,
		Resolver:    fetchApiDestinations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "events"),
		Transform:   transformers.TransformWithStruct(&types.ApiDestination{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ApiDestinationArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApiDestinations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input eventbridge.ListApiDestinationsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEventbridge).Eventbridge
	// No paginator available
	for {
		response, err := svc.ListApiDestinations(ctx, &input, func(options *eventbridge.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ApiDestinations
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
