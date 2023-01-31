// Code generated by codegen; DO NOT EDIT.

package ssm

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Associations() *schema.Table {
	return &schema.Table{
		Name:        "aws_ssm_associations",
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html`,
		Resolver:    fetchSsmAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "association_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "association_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationName"),
			},
			{
				Name:     "association_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AssociationVersion"),
			},
			{
				Name:     "document_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DocumentVersion"),
			},
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceId"),
			},
			{
				Name:     "last_execution_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastExecutionDate"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "overview",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Overview"),
			},
			{
				Name:     "schedule_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduleExpression"),
			},
			{
				Name:     "schedule_offset",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScheduleOffset"),
			},
			{
				Name:     "target_maps",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetMaps"),
			},
			{
				Name:     "targets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Targets"),
			},
		},
	}
}
