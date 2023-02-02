// Code generated by codegen; DO NOT EDIT.

package fsx

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRepositoryTasks() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_data_repository_tasks",
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_DataRepositoryTask.html`,
		Resolver:    fetchFsxDataRepositoryTasks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
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
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Lifecycle"),
			},
			{
				Name:     "task_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskId"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "capacity_to_release",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CapacityToRelease"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndTime"),
			},
			{
				Name:     "failure_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FailureDetails"),
			},
			{
				Name:     "file_cache_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileCacheId"),
			},
			{
				Name:     "file_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemId"),
			},
			{
				Name:     "paths",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Paths"),
			},
			{
				Name:     "report",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Report"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
