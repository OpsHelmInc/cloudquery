// Code generated by codegen; DO NOT EDIT.

package kafka

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:                "aws_kafka_clusters",
		Description:         `https://docs.aws.amazon.com/MSK/2.0/APIReference/v2-clusters-clusterarn.html#v2-clusters-clusterarn-properties`,
		Resolver:            fetchKafkaClusters,
		PreResourceResolver: getCluster,
		Multiplex:           client.ServiceAccountRegionMultiplexer("kafka"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "active_operation_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActiveOperationArn"),
			},
			{
				Name:     "cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterName"),
			},
			{
				Name:     "cluster_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterType"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "current_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CurrentVersion"),
			},
			{
				Name:     "provisioned",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Provisioned"),
			},
			{
				Name:     "serverless",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Serverless"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateInfo"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			Nodes(),
			ClusterOperations(),
		},
	}
}
