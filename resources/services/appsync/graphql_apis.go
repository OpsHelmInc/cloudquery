package appsync

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func GraphqlApis() *schema.Table {
	tableName := "aws_appsync_graphql_apis"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html`,
		Resolver:    fetchAppsyncGraphqlApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appsync"),
		Transform:   transformers.TransformWithStruct(&types.GraphqlApi{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchAppsyncGraphqlApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config appsync.ListGraphqlApisInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppsync).Appsync
	// No paginator available
	for {
		output, err := svc.ListGraphqlApis(ctx, &config, func(options *appsync.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.GraphqlApis
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
