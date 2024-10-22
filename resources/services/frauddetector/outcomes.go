package frauddetector

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Outcomes() *schema.Table {
	tableName := "aws_frauddetector_outcomes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_Outcome.html`,
		Resolver:    fetchFrauddetectorOutcomes,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "frauddetector"),
		Transform:   transformers.TransformWithStruct(&types.Outcome{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveResourceTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchFrauddetectorOutcomes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := frauddetector.NewGetOutcomesPaginator(meta.(*client.Client).Services(client.AWSServiceFrauddetector).Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *frauddetector.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Outcomes
	}
	return nil
}
