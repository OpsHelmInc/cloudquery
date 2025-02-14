package sns

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func Topics() *schema.Table {
	tableName := "aws_sns_topics"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sns/latest/api/API_GetTopicAttributes.html`,
		Resolver:            fetchSnsTopics,
		PreResourceResolver: getTopic,
		Transform:           transformers.TransformWithStruct(&ohaws.Topic{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sns"),
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
				Resolver: resolveSnsTopicTags,
			},
			{
				Name:     "delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("DeliveryPolicy"),
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "effective_delivery_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("EffectiveDeliveryPolicy"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSns).Sns
	config := sns.ListTopicsInput{}
	paginator := sns.NewListTopicsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sns.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Topics
	}
	return nil
}

func getTopic(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSns).Sns
	topic := resource.Item.(types.Topic)

	attrs, err := svc.GetTopicAttributes(ctx,
		&sns.GetTopicAttributesInput{TopicArn: topic.TopicArn},
		func(o *sns.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}

	t := &ohaws.Topic{Arn: topic.TopicArn}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: t})
	if err != nil {
		return err
	}
	if err := dec.Decode(attrs.Attributes); err != nil {
		return err
	}

	resource.Item = t
	return nil
}

func resolveSnsTopicTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	topic := resource.Item.(*ohaws.Topic)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSns).Sns
	tagParams := sns.ListTagsForResourceInput{
		ResourceArn: topic.Arn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams, func(o *sns.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tags.Tags))
}
