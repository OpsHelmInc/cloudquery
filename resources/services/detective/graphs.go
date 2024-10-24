package detective

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Graphs() *schema.Table {
	tableName := "aws_detective_graphs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/detective/latest/APIReference/API_ListGraphs.html`,
		Resolver:    fetchGraphs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.detective"),
		Transform:   transformers.TransformWithStruct(&types.Graph{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGraphTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: schema.Tables{
			members(),
		},
	}
}

func fetchGraphs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDetective).Detective
	config := detective.ListGraphsInput{}
	paginator := detective.NewListGraphsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *detective.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- page.GraphList
	}

	return nil
}

func resolveGraphTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	graph := resource.Item.(types.Graph)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDetective).Detective
	input := &detective.ListTagsForResourceInput{
		ResourceArn: graph.Arn,
	}
	response, err := svc.ListTagsForResource(ctx, input, func(options *detective.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
