package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchFirehoseDeliveryStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Firehose
	input := firehose.ListDeliveryStreamsInput{}
	for {
		response, err := svc.ListDeliveryStreams(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.DeliveryStreamNames
		if !aws.ToBool(response.HasMoreDeliveryStreams) {
			break
		}
		input.ExclusiveStartDeliveryStreamName = aws.String(response.DeliveryStreamNames[len(response.DeliveryStreamNames)-1])
	}
	return nil
}

func getDeliveryStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := c.Services().Firehose
	streamSummary, err := svc.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName: aws.String(streamName),
	})
	if err != nil {
		return err
	}

	tags, err := getFirehoseDeliveryStreamTags(ctx, svc, *streamSummary.DeliveryStreamDescription.DeliveryStreamName)
	if err != nil {
		return err
	}

	resource.Item = &ohaws.FirehoseStream{
		DeliveryStreamDescription: *streamSummary.DeliveryStreamDescription,
		Tags:                      tags,
	}

	return nil
}

func getFirehoseDeliveryStreamTags(ctx context.Context, svc services.FirehoseClient, streamName string) ([]types.Tag, error) {
	input := firehose.ListTagsForDeliveryStreamInput{
		DeliveryStreamName: aws.String(streamName),
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForDeliveryStream(ctx, &input)
		if err != nil {
			return nil, err
		}
		tags = append(tags, output.Tags...)
		if !aws.ToBool(output.HasMoreTags) {
			break
		}
		input.ExclusiveStartTagKey = aws.String(*output.Tags[len(output.Tags)-1].Key)
	}
	return tags, nil
}
