package xray

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Groups() *schema.Table {
	tableName := "aws_xray_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/xray/latest/api/API_Group.html`,
		Resolver:    fetchXrayGroups,
		Transform:   transformers.TransformWithStruct(&types.GroupSummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "xray"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("GroupARN"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveXrayGroupTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchXrayGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := xray.NewGetGroupsPaginator(cl.Services(client.AWSServiceXray).Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *xray.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.Groups
	}
	return nil
}

func resolveXrayGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	group := resource.Item.(types.GroupSummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceXray).Xray
	params := xray.ListTagsForResourceInput{ResourceARN: group.GroupARN}

	output, err := svc.ListTagsForResource(ctx, &params, func(o *xray.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
