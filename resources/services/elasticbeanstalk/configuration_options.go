// Code generated by codegen; DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ConfigurationOptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticbeanstalk_configuration_options",
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationOptionDescription.html`,
		Resolver:    fetchElasticbeanstalkConfigurationOptions,
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
				Name:     "environment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "change_severity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChangeSeverity"),
			},
			{
				Name:     "default_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultValue"),
			},
			{
				Name:     "max_length",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxLength"),
			},
			{
				Name:     "max_value",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxValue"),
			},
			{
				Name:     "min_value",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinValue"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "regex",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Regex"),
			},
			{
				Name:     "user_defined",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("UserDefined"),
			},
			{
				Name:     "value_options",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ValueOptions"),
			},
			{
				Name:     "value_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ValueType"),
			},
			{
				Name:     "application_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplicationArn"),
			},
		},
	}
}
