package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iot"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/services"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func resolveIotTags(ctx context.Context, meta schema.ClientMeta, svc services.IotClient, resource *schema.Resource, c schema.Column, resourceArn *string) error {
	cl := meta.(*client.Client)
	input := iot.ListTagsForResourceInput{
		ResourceArn: resourceArn,
	}
	tags := make(map[string]string)
	paginator := iot.NewListTagsForResourcePaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		client.TagsIntoMap(page.Tags, tags)
	}
	return resource.Set(c.Name, tags)
}
