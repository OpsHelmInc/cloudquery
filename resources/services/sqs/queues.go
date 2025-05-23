package sqs

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func Queues() *schema.Table {
	tableName := "aws_sqs_queues"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueAttributes.html`,
		Resolver:            fetchSqsQueues,
		PreResourceResolver: getQueue,
		Transform:           transformers.TransformWithStruct(&ohaws.Queue{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "sqs"),
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
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "redrive_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedrivePolicy"),
			},
			{
				Name:     "redrive_allow_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("RedriveAllowPolicy"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchSqsQueues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSqs).Sqs
	var params sqs.ListQueuesInput
	paginator := sqs.NewListQueuesPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *sqs.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.QueueUrls
	}
	return nil
}

func getQueue(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSqs).Sqs
	qURL := resource.Item.(string)

	input := sqs.GetQueueAttributesInput{
		QueueUrl:       aws.String(qURL),
		AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll},
	}
	out, err := svc.GetQueueAttributes(ctx, &input, func(o *sqs.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	q := &ohaws.Queue{URL: qURL}
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: q})
	if err != nil {
		return err
	}
	if err := d.Decode(out.Attributes); err != nil {
		return err
	}

	tagsResp, err := svc.ListQueueTags(ctx, &sqs.ListQueueTagsInput{QueueUrl: aws.String(qURL)}, func(o *sqs.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	q.Tags = tagsResp.Tags

	resource.Item = q
	return nil
}
