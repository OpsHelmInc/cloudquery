package route53resolver

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func ResolverQueryLogConfigs() *schema.Table {
	tableName := "aws_route53resolver_resolver_query_log_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverQueryLogConfig.html`,
		Resolver:    fetchQueryLogConfigs,
		Transform:   transformers.TransformWithStruct(&types.ResolverQueryLogConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
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

func fetchQueryLogConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53resolver).Route53resolver
	var input route53resolver.ListResolverQueryLogConfigsInput
	paginator := route53resolver.NewListResolverQueryLogConfigsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ResolverQueryLogConfigs
	}
	return nil
}
