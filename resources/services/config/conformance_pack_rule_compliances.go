// Code generated by codegen; DO NOT EDIT.

package config

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ConformancePackRuleCompliances() *schema.Table {
	return &schema.Table{
		Name:      "aws_config_conformance_pack_rule_compliances",
		Resolver:  fetchConfigConformancePackRuleCompliances,
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
				Name:     "conformance_pack_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "compliance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ComplianceType"),
			},
			{
				Name:     "config_rule_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigRuleName"),
			},
			{
				Name:     "controls",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Controls"),
			},
			{
				Name:     "config_rule_invoked_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ConfigRuleInvokedTime"),
			},
			{
				Name:     "evaluation_result_identifier",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EvaluationResultIdentifier"),
			},
			{
				Name:     "result_recorded_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ResultRecordedTime"),
			},
			{
				Name:     "annotation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Annotation"),
			},
		},
	}
}
