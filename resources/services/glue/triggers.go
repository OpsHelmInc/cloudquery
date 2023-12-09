// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func Triggers() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_triggers",
		Resolver:            fetchGlueTriggers,
		PreResourceResolver: getTrigger,
		Multiplex:           client.ServiceAccountRegionMultiplexer("glue"),
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
				Resolver: resolveGlueTriggerArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueTriggerTags,
			},
			{
				Name:     "actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Actions"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "event_batching_condition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EventBatchingCondition"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "predicate",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Predicate"),
			},
			{
				Name:     "schedule",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Schedule"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "workflow_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkflowName"),
			},
		},
	}
}
