package lightsail

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Databases() *schema.Table {
	tableName := "aws_lightsail_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabase.html`,
		Resolver:    fetchLightsailDatabases,
		Transform:   transformers.TransformWithStruct(&types.RelationalDatabase{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			databaseParameters(),
			databaseEvents(),
			databaseLogEvents(),
		},
	}
}

func fetchLightsailDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetRelationalDatabasesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLightsail).Lightsail
	// No paginator available
	for {
		response, err := svc.GetRelationalDatabases(ctx, &input, func(options *lightsail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.RelationalDatabases
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
