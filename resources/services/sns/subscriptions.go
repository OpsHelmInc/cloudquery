package sns

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func Subscriptions() *schema.Table {
	tableName := "aws_sns_subscriptions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html`,
		Resolver:            fetchSnsSubscriptions,
		PreResourceResolver: getSnsSubscription,
		Transform:           transformers.TransformWithStruct(&ohaws.Subscription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sns"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("SubscriptionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
			{
				Name:     "filter_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("FilterPolicy"),
			},
			{
				Name:     "redrive_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSns).Sns
	config := sns.ListSubscriptionsInput{}
	paginator := sns.NewListSubscriptionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sns.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Subscriptions
	}
	return nil
}

func getSnsSubscription(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSns).Sns
	item := resource.Item.(types.Subscription)
	s := ohaws.Subscription{
		SubscriptionArn: item.SubscriptionArn,
		Owner:           item.Owner,
		Protocol:        item.Protocol,
		TopicArn:        item.TopicArn,
		Endpoint:        item.Endpoint,
	}
	// Return early if SubscriptionARN is not set because it is still pending
	if aws.ToString(item.SubscriptionArn) == "PendingConfirmation" {
		resource.Item = s
		return nil
	}

	attrs, err := svc.GetSubscriptionAttributes(ctx,
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn},
		func(o *sns.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		// If a subscriptions topic is deleted GetSubscriptionAttributes will error.
		if client.IsAWSError(err, "NotFound") {
			resource.Item = s
			return nil
		}
		return err
	}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &s})
	if err != nil {
		return err
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		return err
	}
	resource.Item = s
	return nil
}
