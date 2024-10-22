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

func Connections() *schema.Table {
	tableName := "aws_apprunner_connections"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_Connection.html`,
		Resolver:    fetchApprunnerConnections,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:   transformers.TransformWithStruct(&types.ConnectionSummary{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ConnectionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveApprunnerTags("ConnectionArn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApprunnerConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListConnectionsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApprunner).Apprunner
	paginator := apprunner.NewListConnectionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.ConnectionSummaryList
	}
	return nil
}
