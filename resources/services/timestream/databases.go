package timestream

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Databases() *schema.Table {
	tableName := "aws_timestream_databases"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/timestream/latest/developerguide/API_Database.html`,
		Resolver:    fetchTimestreamDatabases,
		Transform:   transformers.TransformWithStruct(&types.Database{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ingest.timestream"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: fetchDatabaseTags,
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
			tables(),
		},
	}
}

func fetchTimestreamDatabases(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	services := cl.Services(client.AWSServiceTimestreamwrite)
	svc := services.Timestreamwrite
	input := &timestreamwrite.ListDatabasesInput{MaxResults: aws.Int32(20)}
	paginator := timestreamwrite.NewListDatabasesPaginator(svc, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(o *timestreamwrite.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Databases
	}
	return nil
}

func fetchDatabaseTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceTimestreamwrite).Timestreamwrite

	output, err := svc.ListTagsForResource(ctx,
		&timestreamwrite.ListTagsForResourceInput{
			ResourceARN: resource.Item.(types.Database).Arn,
		},
		func(o *timestreamwrite.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
