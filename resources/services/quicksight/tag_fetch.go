package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/thoas/go-funk"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

var tagsCol = schema.Column{
	Name:     "tags",
	Type:     sdkTypes.ExtensionTypes.JSON,
	Resolver: resolveTags,
}

func resolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	arn := funk.Get(r.Item, "Arn", funk.WithAllowZero()).(*string)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	params := quicksight.ListTagsForResourceInput{ResourceArn: arn}

	output, err := svc.ListTagsForResource(ctx, &params, func(options *quicksight.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return r.Set(c.Name, client.TagsToMap(output.Tags))
}
