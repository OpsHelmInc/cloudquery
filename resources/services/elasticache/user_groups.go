// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func UserGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_user_groups",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UserGroup.html`,
		Resolver:    fetchElasticacheUserGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "pending_changes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingChanges"),
			},
			{
				Name:     "replication_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReplicationGroups"),
			},
			{
				Name:     "serverless_caches",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ServerlessCaches"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "user_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserGroupId"),
			},
			{
				Name:     "user_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("UserIds"),
			},
		},
	}
}
