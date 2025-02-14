package kinesis

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/services"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/ohaws"
)

func Streams() *schema.Table {
	tableName := "aws_kinesis_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Transform:           transformers.TransformWithStruct(&ohaws.KinesisStream{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "kinesis"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("StreamARN"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchKinesisStreams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceKinesis).Kinesis
	input := kinesis.ListStreamsInput{}
	paginator := kinesis.NewListStreamsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *kinesis.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.StreamNames
	}
	return nil
}

func getStream(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := c.Services().Kinesis
	streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	}, func(o *kinesis.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}

	tags, err := getKinesisStreamTags(ctx, meta, svc, streamName)
	if err != nil {
		return err
	}

	resource.Item = &ohaws.KinesisStream{
		StreamDescriptionSummary: *streamSummary.StreamDescriptionSummary,
		Tags:                     tags,
	}
	return nil
}

func getKinesisStreamTags(ctx context.Context, meta schema.ClientMeta, svc services.KinesisClient, streamName string) ([]types.Tag, error) {
	c := meta.(*client.Client)
	input := kinesis.ListTagsForStreamInput{
		StreamName: aws.String(streamName),
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForStream(ctx, &input, func(o *kinesis.Options) {
			o.Region = c.Region
		})
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
