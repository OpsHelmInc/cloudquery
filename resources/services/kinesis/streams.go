package kinesis

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
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
	resource.Item = streamSummary.StreamDescriptionSummary
	return nil
}

func resolveKinesisStreamTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceKinesis).Kinesis
	summary := resource.Item.(*types.StreamDescriptionSummary)
	input := kinesis.ListTagsForStreamInput{
		StreamName: summary.StreamName,
	}
	var tags []types.Tag
	for {
		output, err := svc.ListTagsForStream(ctx, &input, func(options *kinesis.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, output.Tags...)
		if !aws.ToBool(output.HasMoreTags) {
			break
		}
		input.ExclusiveStartTagKey = aws.String(*output.Tags[len(output.Tags)-1].Key)
	}
	return resource.Set(c.Name, client.TagsToMap(tags))
}
