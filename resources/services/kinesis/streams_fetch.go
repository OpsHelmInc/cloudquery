package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"

	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"

	"github.com/OpsHelmInc/ohaws"
)

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Kinesis
	input := kinesis.ListStreamsInput{}
	for {
		response, err := svc.ListStreams(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.StreamNames
		if !aws.ToBool(response.HasMoreStreams) {
			break
		}
		input.ExclusiveStartStreamName = aws.String(response.StreamNames[len(response.StreamNames)-1])
	}
	return nil
}

func getStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := c.Services().Kinesis
	streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		return err
	}

	tags, err := getKinesisStreamTags(ctx, svc, streamName)
	if err != nil {
		return err
	}

	resource.Item = &ohaws.KinesisStream{
		StreamDescriptionSummary: *streamSummary.StreamDescriptionSummary,
		Tags:                     tags,
	}
	return nil
}

func getKinesisStreamTags(ctx context.Context, svc services.KinesisClient, streamName string) ([]types.Tag, error) {
	input := kinesis.ListTagsForStreamInput{
		StreamName: aws.String(streamName),
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForStream(ctx, &input)
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

func resolveKinesisStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*ohaws.KinesisStream)
	return resource.Set(c.Name, client.TagsToMap(r.Tags))
}
