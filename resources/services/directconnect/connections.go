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

func Connections() *schema.Table {
	tableName := "aws_directconnect_connections"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_Connection.html`,
		Resolver:    fetchDirectconnectConnections,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.Connection{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveConnectionARN(),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ConnectionId"),
				PrimaryKeyComponent: true,
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

func fetchDirectconnectConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeConnectionsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDirectconnect).Directconnect
	output, err := svc.DescribeConnections(ctx, &config, func(options *directconnect.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.Connections
	return nil
}

func resolveConnectionARN() schema.ColumnResolver {
	return client.ResolveARN(client.DirectConnectService, func(resource *schema.Resource) ([]string, error) {
		return []string{"dxcon", *resource.Item.(types.Connection).ConnectionId}, nil
	})
}
