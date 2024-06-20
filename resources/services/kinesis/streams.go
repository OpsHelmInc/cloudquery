package kinesis

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Streams() *schema.Table {
	tableName := "aws_kinesis_streams"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Transform:           transformers.TransformWithStruct(&types.StreamDescriptionSummary{}),
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
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveKinesisStreamTags,
			},
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
	cl := meta.(*client.Client)
	streamName := resource.Item.(string)
	svc := cl.Services(client.AWSServiceKinesis).Kinesis
	streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
		StreamName: aws.String(streamName),
	}, func(options *kinesis.Options) {
		options.Region = cl.Region
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
