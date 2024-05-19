package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchSnsSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	config := sns.ListSubscriptionsInput{}
	paginator := sns.NewListSubscriptionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sns.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.Subscriptions
	}
	return nil
}

func getSnsSubscription(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
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

	attrs, err := svc.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{SubscriptionArn: item.SubscriptionArn})
	if err != nil {
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
