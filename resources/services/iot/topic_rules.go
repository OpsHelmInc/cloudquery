// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TopicRules() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_topic_rules",
		Resolver:  fetchIotTopicRules,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: ResolveIotTopicRuleTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RuleArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Rule"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
