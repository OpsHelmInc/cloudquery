package athena

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func workGroupNamedQueries() *schema.Table {
	tableName := "aws_athena_work_group_named_queries"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_NamedQuery.html`,
		Resolver:            fetchAthenaWorkGroupNamedQueries,
		PreResourceResolver: getWorkGroupNamedQuery,
		Transform:           transformers.TransformWithStruct(&types.NamedQuery{}, transformers.WithPrimaryKeyComponents("NamedQueryId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "work_group_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAthenaWorkGroupNamedQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAthena).Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListNamedQueriesInput{WorkGroup: wg.Name}
	paginator := athena.NewListNamedQueriesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *athena.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.NamedQueryIds
	}
	return nil
}

func getWorkGroupNamedQuery(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAthena).Athena

	d := resource.Item.(string)
	dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
		NamedQueryId: aws.String(d),
	}, func(options *athena.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.NamedQuery
	return nil
}
