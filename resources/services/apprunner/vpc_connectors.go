package apprunner

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func VpcConnectors() *schema.Table {
	tableName := "aws_apprunner_vpc_connectors"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html`,
		Resolver:    fetchApprunnerVpcConnectors,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:   transformers.TransformWithStruct(&types.VpcConnector{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("VpcConnectorArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveApprunnerTags("VpcConnectorArn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApprunnerVpcConnectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListVpcConnectorsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApprunner).Apprunner
	paginator := apprunner.NewListVpcConnectorsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.VpcConnectors
	}
	return nil
}
