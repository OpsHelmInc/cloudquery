package quicksight

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/thoas/go-funk"
)

func resolveTags() schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arn := funk.Get(r.Item, "Arn", funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services().Quicksight
		params := quicksight.ListTagsForResourceInput{ResourceArn: arn}

		output, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
