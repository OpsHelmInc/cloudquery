// Code generated by codegen; DO NOT EDIT.

package backup

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PlanSelections() *schema.Table {
	return &schema.Table{
		Name:        "aws_backup_plan_selections",
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html`,
		Resolver:    fetchBackupPlanSelections,
		Multiplex:   client.ServiceAccountRegionMultiplexer("backup"),
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
				Name:     "plan_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "backup_plan_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupPlanId"),
			},
			{
				Name:     "backup_selection",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackupSelection"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "creator_request_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatorRequestId"),
			},
			{
				Name:     "selection_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelectionId"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
