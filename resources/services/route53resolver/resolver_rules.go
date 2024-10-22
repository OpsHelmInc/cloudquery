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

func ResolverRules() *schema.Table {
	tableName := "aws_route53resolver_resolver_rules"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverRule.html`,
		Resolver:    fetchResolverRules,
		Transform:   transformers.TransformWithStruct(&types.ResolverRule{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
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

func fetchResolverRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53resolver).Route53resolver
	var input route53resolver.ListResolverRulesInput
	paginator := route53resolver.NewListResolverRulesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ResolverRules
	}
	return nil
}
