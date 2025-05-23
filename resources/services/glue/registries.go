package glue

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Registries() *schema.Table {
	tableName := "aws_glue_registries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_RegistryListItem.html`,
		Resolver:    fetchGlueRegistries,
		Transform:   transformers.TransformWithStruct(&types.RegistryListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("RegistryArn"),
				PrimaryKeyComponent: true,
			},
			tagsCol(func(_ *client.Client, resource *schema.Resource) string {
				return *resource.Item.(types.RegistryListItem).RegistryArn
			}),
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			registrySchemas(),
		},
	}
}

func fetchGlueRegistries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	paginator := glue.NewListRegistriesPaginator(svc, &glue.ListRegistriesInput{
		MaxResults: aws.Int32(100),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Registries
	}
	return nil
}
