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

func FirewallRuleGroups() *schema.Table {
	tableName := "aws_route53resolver_firewall_rule_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_FirewallRuleGroup.html`,
		Resolver:            fetchFirewallRuleGroups,
		PreResourceResolver: getFirewallRuleGroups,
		Transform:           transformers.TransformWithStruct(&types.FirewallRuleGroup{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "route53resolver"),
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

func fetchFirewallRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53resolver).Route53resolver
	var input route53resolver.ListFirewallRuleGroupsInput
	paginator := route53resolver.NewListFirewallRuleGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *route53resolver.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.FirewallRuleGroups
	}
	return nil
}

func getFirewallRuleGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53resolver).Route53resolver
	v := resource.Item.(types.FirewallRuleGroupMetadata)

	d, err := svc.GetFirewallRuleGroup(ctx, &route53resolver.GetFirewallRuleGroupInput{FirewallRuleGroupId: v.Id}, func(options *route53resolver.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = d.FirewallRuleGroup

	return nil
}
