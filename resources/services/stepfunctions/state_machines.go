// Code generated by codegen; DO NOT EDIT.

package stepfunctions

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StateMachines() *schema.Table {
	return &schema.Table{
		Name:                "aws_stepfunctions_state_machines",
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html`,
		Resolver:            fetchStepfunctionsStateMachines,
		PreResourceResolver: getStepFunction,
		Multiplex:           client.ServiceAccountRegionMultiplexer("states"),
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
				Resolver: schema.PathResolver("StateMachineArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveStepFunctionTags,
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::StepFunctions::StateMachine"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "definition",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Definition"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleArn"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "label",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Label"),
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoggingConfiguration"),
			},
			{
				Name:     "revision_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevisionId"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "tracing_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TracingConfiguration"),
			},
		},
	}
}
