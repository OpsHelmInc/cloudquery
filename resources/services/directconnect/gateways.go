package directconnect

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Gateways() *schema.Table {
	tableName := "aws_directconnect_gateways"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html`,
		Resolver:    fetchDirectconnectGateways,
		Multiplex:   client.AccountMultiplex(tableName),
		Transform:   transformers.TransformWithStruct(&types.DirectConnectGateway{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGatewayARN,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			gatewayAssociations(),
			gatewayAttachments(),
		},
	}
}

func fetchDirectconnectGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config directconnect.DescribeDirectConnectGatewaysInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDirectconnect).Directconnect
	// No paginator available
	for {
		output, err := svc.DescribeDirectConnectGateways(ctx, &config, func(options *directconnect.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.DirectConnectGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveGatewayARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	gw := resource.Item.(types.DirectConnectGateway)

	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "directconnect",
		AccountID: *gw.OwnerAccount,
		Resource:  strings.Join([]string{"dx-gateway", *gw.DirectConnectGatewayId}, "/"),
	}.String())
}
