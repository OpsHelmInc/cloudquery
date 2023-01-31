// Code generated by codegen; DO NOT EDIT.

package wafregional

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RuleGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafregional_rule_groups",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RuleGroup.html`,
		Resolver:    fetchWafregionalRuleGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("waf-regional"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveWafregionalRuleGroupArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveWafregionalRuleGroupTags,
				Description: `Rule group tags.`,
			},
			{
				Name:     "rule_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleGroupId"),
			},
			{
				Name:     "metric_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricName"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
		},
	}
}
