package cloudwatchlogs

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudwatchlogsLogGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatchlogs.DescribeLogGroupsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatchlogs
	for {
		var output []ohaws.LogGroup
		response, err := svc.DescribeLogGroups(ctx, &config)
		if err != nil {
			return err
		}

		for _, group := range response.LogGroups {
			entry := ohaws.LogGroup{
				LogGroup: &group,
			}
			tagRes, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{
				LogGroupName: group.LogGroupName,
			})
			if err != nil {
				return err
			}

			for key, value := range tagRes.Tags {
				entry.Tags = append(entry.Tags, ohaws.Tag{
					Key:   &key,
					Value: &value,
				})
			}

			output = append(output, entry)
		}

		res <- output
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}

func resolveLogGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	lg := resource.Item.(types.LogGroup)
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatchlogs
	out, err := svc.ListTagsLogGroup(ctx, &cloudwatchlogs.ListTagsLogGroupInput{LogGroupName: lg.LogGroupName})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
