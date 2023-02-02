// Code generated by codegen; DO NOT EDIT.

package wafregional

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RateBasedRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafregional_rate_based_rules",
		Description: `https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_RateBasedRule.html`,
		Resolver:    fetchWafregionalRateBasedRules,
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
				Resolver: resolveWafregionalRateBasedRuleArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafregionalRateBasedRuleTags,
			},
			{
				Name:     "match_predicates",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MatchPredicates"),
			},
			{
				Name:     "rate_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RateKey"),
			},
			{
				Name:     "rate_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RateLimit"),
			},
			{
				Name:     "rule_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleId"),
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
