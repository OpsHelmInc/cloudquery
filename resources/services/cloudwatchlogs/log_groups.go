package cloudwatchlogs

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func LogGroups() *schema.Table {
	tableName := "aws_cloudwatchlogs_log_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html`,
		Resolver:    fetchCloudwatchlogsLogGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "logs"),
		Transform:   transformers.TransformWithStruct(&types.LogGroup{}),
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
				Resolver: resolveLogGroupTags,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			subscriptionFilters(),
			dataProtectionPolicy(),
		},
	}
}

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudwatchlogs).Cloudwatchlogs
	paginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudwatchlogs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LogGroups
	}
	return nil
}

func resolveLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudwatchlogs).Cloudwatchlogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName}, func(options *cloudwatchlogs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
