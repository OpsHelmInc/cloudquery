// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Rules() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_rules",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html`,
		Resolver:    fetchFrauddetectorRules,
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "detector_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DetectorId"),
			},
			{
				Name:     "expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Expression"),
			},
			{
				Name:     "language",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Language"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "outcomes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Outcomes"),
			},
			{
				Name:     "rule_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleId"),
			},
			{
				Name:     "rule_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleVersion"),
			},
		},
	}
}
