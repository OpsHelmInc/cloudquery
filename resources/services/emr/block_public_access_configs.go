// Code generated by codegen; DO NOT EDIT.

package emr

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func BlockPublicAccessConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_emr_block_public_access_configs",
		Resolver:  fetchEmrBlockPublicAccessConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
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
				Name:     "block_public_access_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockPublicAccessConfiguration"),
			},
			{
				Name:     "block_public_access_configuration_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockPublicAccessConfigurationMetadata"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
