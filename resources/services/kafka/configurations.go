// Code generated by codegen; DO NOT EDIT.

package kafka

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Configurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_kafka_configurations",
		Description: `https://docs.aws.amazon.com/msk/1.0/apireference/clusters-clusterarn-configuration.html`,
		Resolver:    fetchKafkaConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("kafka"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
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
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::MSK::Configuration"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "kafka_versions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("KafkaVersions"),
			},
			{
				Name:     "latest_revision",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LatestRevision"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
		},
	}
}
