package qldb

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Ledgers() *schema.Table {
	tableName := "aws_qldb_ledgers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/qldb/latest/developerguide/API_DescribeLedger.html`,
		Resolver:            fetchQldbLedgers,
		PreResourceResolver: getLedger,
		Transform:           transformers.TransformWithStruct(&qldb.DescribeLedgerOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "qldb"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:        "tags",
				Type:        sdkTypes.ExtensionTypes.JSON,
				Resolver:    resolveQldbLedgerTags,
				Description: `The tags associated with the pipeline.`,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			ledgerJournalKinesisStreams(),
			ledgerJournalS3Exports(),
		},
	}
}

func fetchQldbLedgers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQldb).Qldb
	config := qldb.ListLedgersInput{}
	paginator := qldb.NewListLedgersPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *qldb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Ledgers
	}
	return nil
}

func getLedger(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQldb).Qldb
	l := resource.Item.(types.LedgerSummary)

	response, err := svc.DescribeLedger(ctx, &qldb.DescribeLedgerInput{Name: l.Name}, func(options *qldb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response
	return nil
}

func resolveQldbLedgerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	ledger := resource.Item.(*qldb.DescribeLedgerOutput)

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQldb).Qldb
	response, err := svc.ListTagsForResource(ctx, &qldb.ListTagsForResourceInput{
		ResourceArn: ledger.Arn,
	}, func(options *qldb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Tags)
}
