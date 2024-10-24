package mq

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/mq"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func brokerUsers() *schema.Table {
	tableName := "aws_mq_broker_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/brokers-broker-id-users-username.html`,
		Resolver:    fetchMqBrokerUsers,
		Transform:   transformers.TransformWithStruct(&mq.DescribeUserOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithPrimaryKeyComponents("Username")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "broker_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceMq).Mq
	for _, us := range broker.Users {
		input := mq.DescribeUserInput{
			BrokerId: broker.BrokerId,
			Username: us.Username,
		}
		output, err := svc.DescribeUser(ctx, &input, func(options *mq.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}
