// Code generated by codegen; DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Applications() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_applications",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationDescription.html`,
		Resolver:    fetchElasticbeanstalkApplications,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticbeanstalk"),
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
				Resolver: schema.PathResolver("ApplicationArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "date_created",
				Type: schema.TypeTimestamp,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::ElasticBeanstalk::Application"),
			},
			{
				Name:     "application_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationName"),
			},
			{
				Name:     "configuration_templates",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ConfigurationTemplates"),
			},
			{
				Name:     "date_updated",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DateUpdated"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "resource_lifecycle_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceLifecycleConfig"),
			},
			{
				Name:     "versions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Versions"),
			},
		},
	}
}
