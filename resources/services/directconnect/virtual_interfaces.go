package directconnect

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func VirtualInterfaces() *schema.Table {
	tableName := "aws_directconnect_virtual_interfaces"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualInterface.html`,
		Resolver:    fetchDirectconnectVirtualInterfaces,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.VirtualInterface{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveVirtualInterfaceARN(),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("VirtualInterfaceId"),
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchDirectconnectVirtualInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeVirtualInterfacesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDirectconnect).Directconnect
	output, err := svc.DescribeVirtualInterfaces(ctx, &config, func(options *directconnect.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.VirtualInterfaces
	return nil
}

func resolveVirtualInterfaceARN() schema.ColumnResolver {
	return client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
		return []string{"dxvif", *resource.Item.(types.VirtualInterface).VirtualInterfaceId}, nil
	})
}
