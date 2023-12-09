// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ClusterParameterGroupParameters() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_parameter_group_parameters",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Parameter.html`,
		Resolver:    fetchRdsClusterParameterGroupParameters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("rds"),
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
				Name:     "cluster_parameter_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "allowed_values",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AllowedValues"),
			},
			{
				Name:     "apply_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyMethod"),
			},
			{
				Name:     "apply_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApplyType"),
			},
			{
				Name:     "data_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "is_modifiable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsModifiable"),
			},
			{
				Name:     "minimum_engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MinimumEngineVersion"),
			},
			{
				Name:     "parameter_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterName"),
			},
			{
				Name:     "parameter_value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParameterValue"),
			},
			{
				Name:     "source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Source"),
			},
			{
				Name:     "supported_engine_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedEngineModes"),
			},
		},
	}
}
