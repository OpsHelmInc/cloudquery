// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func MlTransformTaskRuns() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_ml_transform_task_runs",
		Resolver:  fetchGlueMlTransformTaskRuns,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:     "ml_transform_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "completed_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletedOn"),
			},
			{
				Name:     "error_string",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorString"),
			},
			{
				Name:     "execution_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ExecutionTime"),
			},
			{
				Name:     "last_modified_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedOn"),
			},
			{
				Name:     "log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupName"),
			},
			{
				Name:     "properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Properties"),
			},
			{
				Name:     "started_on",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartedOn"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "task_run_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TaskRunId"),
			},
			{
				Name:     "transform_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransformId"),
			},
		},
	}
}
