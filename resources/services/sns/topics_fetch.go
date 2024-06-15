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

func fetchSnsTopics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	config := sns.ListTopicsInput{}
	for {
		output, err := svc.ListTopics(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Topics

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getTopic(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sns
	topic := resource.Item.(types.Topic)

	attrs, err := svc.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{TopicArn: topic.TopicArn})
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

	tagParams := sns.ListTagsForResourceInput{
		ResourceArn: topic.TopicArn,
	}
	tags, err := svc.ListTagsForResource(ctx, &tagParams)
	if err != nil {
		return err
	}
	t.Tags = tags.Tags

	resource.Item = t
	return nil
}
