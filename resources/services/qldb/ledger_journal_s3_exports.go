package qldb

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func ledgerJournalS3Exports() *schema.Table {
	tableName := "aws_qldb_ledger_journal_s3_exports"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/qldb/latest/developerguide/API_JournalS3ExportDescription.html`,
		Resolver:    fetchQldbLedgerJournalS3Exports,
		Transform:   transformers.TransformWithStruct(&types.JournalS3ExportDescription{}, transformers.WithPrimaryKeyComponents("ExportId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "ledger_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchQldbLedgerJournalS3Exports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	ledger := parent.Item.(*qldb.DescribeLedgerOutput)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQldb).Qldb
	config := &qldb.ListJournalS3ExportsForLedgerInput{
		Name:       ledger.Name,
		MaxResults: aws.Int32(100),
	}
	paginator := qldb.NewListJournalS3ExportsForLedgerPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *qldb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- page.JournalS3Exports
	}
	return nil
}
