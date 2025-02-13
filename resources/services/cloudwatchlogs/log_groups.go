// Code generated by codegen; DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LogGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_log_groups",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html`,
		Resolver:    fetchCloudwatchlogsLogGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveLogGroupTags,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "data_protection_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataProtectionStatus"),
			},
			{
				Name:     "inherited_properties",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("InheritedProperties"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "log_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupArn"),
			},
			{
				Name:     "log_group_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupClass"),
			},
			{
				Name:     "log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupName"),
			},
			{
				Name:     "metric_filter_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MetricFilterCount"),
			},
			{
				Name:     "retention_in_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RetentionInDays"),
			},
			{
				Name:     "stored_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StoredBytes"),
			},
		},
	}
}
