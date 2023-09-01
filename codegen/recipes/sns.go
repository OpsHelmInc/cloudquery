package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/resources/services/sns/models"
)

func SNSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "subscriptions",
			Struct:              &models.Subscription{},
			SkipFields:          []string{"SubscriptionArn", "DeliveryPolicy", "EffectiveDeliveryPolicy", "FilterPolicy", "RedrivePolicy"},
			PreResourceResolver: "getSnsSubscription",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("SubscriptionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "delivery_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("DeliveryPolicy")`,
					},
					{
						Name:     "effective_delivery_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("EffectiveDeliveryPolicy")`,
					},
					{
						Name:     "filter_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("FilterPolicy")`,
					},
					{
						Name:     "redrive_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("RedrivePolicy")`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::SNS::Subscription")`,
					},
				}...),
		},

		{
			SubService:          "topics",
			Struct:              &models.Topic{},
			SkipFields:          []string{"Arn", "Policy", "EffectiveDeliveryPolicy", "DeliveryPolicy"},
			PreResourceResolver: "getTopic",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveSnsTopicTags`,
					},
					{
						Name:     "delivery_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("DeliveryPolicy")`,
					},
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("Policy")`,
					},
					{
						Name:     "effective_delivery_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("EffectiveDeliveryPolicy")`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::SNS::Topic")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "sns"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("sns")`
	}
	return resources
}
