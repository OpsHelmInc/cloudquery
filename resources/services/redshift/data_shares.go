package redshift

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func DataShares() *schema.Table {
	tableName := "aws_redshift_data_shares"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_DataShare.html`,
		Resolver:    fetchDataShares,
		Transform:   transformers.TransformWithStruct(&types.DataShare{}, transformers.WithPrimaryKeyComponents("ProducerArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "redshift"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DataShareArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchDataShares(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRedshift).Redshift

	config := redshift.DescribeDataSharesInput{
		MaxRecords: aws.Int32(100),
	}
	paginator := redshift.NewDescribeDataSharesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *redshift.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DataShares
	}

	return nil
}
