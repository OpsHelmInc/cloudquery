package networkmanager

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func links() *schema.Table {
	tableName := "aws_networkmanager_links"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_Link.html
The  'request_region' column is added to show region of where the request was made from.`,
		Resolver:  fetchLinks,
		Transform: transformers.TransformWithStruct(&types.Link{}, transformers.WithPrimaryKeyComponents("GlobalNetworkId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.RequestRegionColumn(true), {
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("LinkArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: schema.Tables{},
	}
}

func fetchLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkmanager).Networkmanager
	globalNetwork := parent.Item.(types.GlobalNetwork)
	input := &networkmanager.GetLinksInput{
		GlobalNetworkId: globalNetwork.GlobalNetworkId,
	}
	paginator := networkmanager.NewGetLinksPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *networkmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Links
	}
	return nil
}
