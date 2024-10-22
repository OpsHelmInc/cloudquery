package ram

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ResourceShares() *schema.Table {
	tableName := "aws_ram_resource_shares"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShare.html`,
		Resolver:    fetchRamResourceShares,
		Transform:   transformers.TransformWithStruct(&types.ResourceShare{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ResourceShareArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			resourceSharePermissions(),
		},
	}
}

func fetchRamResourceShares(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	err := fetchRamResourceSharesByType(ctx, meta, types.ResourceOwnerOtherAccounts, res)
	if err != nil {
		return err
	}
	err = fetchRamResourceSharesByType(ctx, meta, types.ResourceOwnerSelf, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamResourceSharesByType(ctx context.Context, meta schema.ClientMeta, shareType types.ResourceOwner, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &ram.GetResourceSharesInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: shareType,
	}
	paginator := ram.NewGetResourceSharesPaginator(meta.(*client.Client).Services(client.AWSServiceRam).Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ResourceShares
	}
	return nil
}
