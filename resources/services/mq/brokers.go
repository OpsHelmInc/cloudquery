package mq

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Brokers() *schema.Table {
	tableName := "aws_mq_brokers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers.html`,
		Resolver:            fetchMqBrokers,
		PreResourceResolver: getMqBroker,
		Transform:           transformers.TransformWithStruct(&mq.DescribeBrokerOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("BrokerArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			brokerConfigurations(),
			brokerUsers(),
		},
	}
}

func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config mq.ListBrokersInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceMq).Mq
	paginator := mq.NewListBrokersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *mq.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.BrokerSummaries
	}
	return nil
}

func getMqBroker(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceMq).Mq
	bs := resource.Item.(types.BrokerSummary)

	output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, func(options *mq.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
