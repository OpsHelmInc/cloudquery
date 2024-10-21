package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/thoas/go-funk"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func resolveRDSTags(path string) schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, path, funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services(client.AWSServiceRds).Rds
		input := rds.ListTagsForResourceInput{ResourceName: arn}
		output, err := svc.ListTagsForResource(ctx, &input, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.TagList))
	}
}
