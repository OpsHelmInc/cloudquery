// Code generated by codegen; DO NOT EDIT.

package config

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ConfigRuleCompliances() *schema.Table {
	return &schema.Table{
		Name:      "aws_config_config_rule_compliances",
		Resolver:  fetchConfigConfigRuleCompliances,
		Multiplex: client.ServiceAccountRegionMultiplexer("config"),
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
				Name:     "compliance",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Compliance"),
			},
			{
				Name:     "config_rule_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigRuleName"),
			},
		},
	}
}
