// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BatchImports() *schema.Table {
	return &schema.Table{
		Name:        "aws_frauddetector_batch_imports",
		Description: `https://docs.aws.amazon.com/frauddetector/latest/api/API_BatchImport.html`,
		Resolver:    fetchFrauddetectorBatchImports,
		Multiplex:   client.ServiceAccountRegionMultiplexer("frauddetector"),
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
				Name:     "completion_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CompletionTime"),
			},
			{
				Name:     "event_type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTypeName"),
			},
			{
				Name:     "failed_records_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("FailedRecordsCount"),
			},
			{
				Name:     "failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureReason"),
			},
			{
				Name:     "iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleArn"),
			},
			{
				Name:     "input_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InputPath"),
			},
			{
				Name:     "job_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("JobId"),
			},
			{
				Name:     "output_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutputPath"),
			},
			{
				Name:     "processed_records_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ProcessedRecordsCount"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "total_records_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalRecordsCount"),
			},
		},
	}
}
