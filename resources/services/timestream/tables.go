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
)

func tables() *schema.Table {
	return &schema.Table{
		Name:        "aws_timestream_tables",
		Description: `https://docs.aws.amazon.com/timestream/latest/developerguide/API_Table.html`,
		Resolver:    fetchTimestreamTables,
		Transform:   transformers.TransformWithStruct(&types.Table{}),
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

func fetchTimestreamTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &timestreamwrite.ListTablesInput{
		DatabaseName: parent.Item.(types.Database).DatabaseName,
		MaxResults:   aws.Int32(20),
	}
	paginator := timestreamwrite.NewListTablesPaginator(cl.Services(client.AWSServiceTimestreamwrite).Timestreamwrite, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(o *timestreamwrite.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Tables
	}
	return nil
}
