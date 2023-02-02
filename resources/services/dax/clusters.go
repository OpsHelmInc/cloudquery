// Code generated by codegen; DO NOT EDIT.

package dax

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_dax_clusters",
		Description: `https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_dax_Cluster.html`,
		Resolver:    fetchDaxClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("dax"),
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
				Resolver: schema.PathResolver("ClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveClusterTags,
			},
			{
				Name:     "active_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ActiveNodes"),
			},
			{
				Name:     "cluster_discovery_endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterDiscoveryEndpoint"),
			},
			{
				Name:     "cluster_endpoint_encryption_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterEndpointEncryptionType"),
			},
			{
				Name:     "cluster_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleArn"),
			},
			{
				Name:     "node_ids_to_remove",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NodeIdsToRemove"),
			},
			{
				Name:     "node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeType"),
			},
			{
				Name:     "nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Nodes"),
			},
			{
				Name:     "notification_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NotificationConfiguration"),
			},
			{
				Name:     "parameter_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ParameterGroup"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "sse_description",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SSEDescription"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subnet_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetGroup"),
			},
			{
				Name:     "total_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalNodes"),
			},
		},
	}
}
