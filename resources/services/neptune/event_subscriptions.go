package neptune

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func EventSubscriptions() *schema.Table {
	tableName := "aws_neptune_event_subscriptions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-events.html#DescribeEventSubscriptions`,
		Resolver:    fetchNeptuneEventSubscriptions,
		Transform:   transformers.TransformWithStruct(&types.EventSubscription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("EventSubscriptionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveNeptuneEventSubscriptionTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchNeptuneEventSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNeptune).Neptune
	input := neptune.DescribeEventSubscriptionsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
	paginator := neptune.NewDescribeEventSubscriptionsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *neptune.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.EventSubscriptionsList
	}
	return nil
}

func resolveNeptuneEventSubscriptionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(types.EventSubscription)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNeptune).Neptune
	out, err := svc.ListTagsForResource(ctx, &neptune.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn}, func(options *neptune.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.TagList))
}
