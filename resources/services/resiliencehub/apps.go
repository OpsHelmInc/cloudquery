package resiliencehub

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Apps() *schema.Table {
	tableName := "aws_resiliencehub_apps"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_App.html`,
		Resolver:            fetchApps,
		PreResourceResolver: describeApp,
		Transform:           transformers.TransformWithStruct(&types.App{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("AppArn"),
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{appAssesments(), appVersions()},
	}
}

func fetchApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	p := resiliencehub.NewListAppsPaginator(svc, &resiliencehub.ListAppsInput{})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.AppSummaries
	}
	return nil
}

func describeApp(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	out, err := svc.DescribeApp(ctx,
		&resiliencehub.DescribeAppInput{AppArn: resource.Item.(types.AppSummary).AppArn},
		func(options *resiliencehub.Options) {
			options.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	resource.SetItem(out.App)
	return nil
}
