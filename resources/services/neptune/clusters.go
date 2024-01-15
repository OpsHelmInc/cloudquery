// Code generated by codegen; DO NOT EDIT.

package neptune

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_neptune_clusters",
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-clusters.html#DescribeDBClusters`,
		Resolver:    fetchNeptuneClusters,
		Multiplex:   client.ServiceAccountRegionMultiplexer("neptune"),
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
				Resolver: schema.PathResolver("DBClusterArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveNeptuneClusterTags,
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::Neptune::DBCluster"),
			},
			{
				Name:     "allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedStorage"),
			},
			{
				Name:     "associated_roles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AssociatedRoles"),
			},
			{
				Name:     "automatic_restart_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AutomaticRestartTime"),
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "backup_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BackupRetentionPeriod"),
			},
			{
				Name:     "character_set_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CharacterSetName"),
			},
			{
				Name:     "clone_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloneGroupId"),
			},
			{
				Name:     "cluster_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClusterCreateTime"),
			},
			{
				Name:     "copy_tags_to_snapshot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CopyTagsToSnapshot"),
			},
			{
				Name:     "cross_account_clone",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CrossAccountClone"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_members",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBClusterMembers"),
			},
			{
				Name:     "db_cluster_option_group_memberships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBClusterOptionGroupMemberships"),
			},
			{
				Name:     "db_cluster_parameter_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterParameterGroup"),
			},
			{
				Name:     "db_subnet_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name:     "database_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseName"),
			},
			{
				Name:     "db_cluster_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DbClusterResourceId"),
			},
			{
				Name:     "deletion_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DeletionProtection"),
			},
			{
				Name:     "earliest_restorable_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EarliestRestorableTime"),
			},
			{
				Name:     "enabled_cloudwatch_logs_exports",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledCloudwatchLogsExports"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "global_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalClusterIdentifier"),
			},
			{
				Name:     "hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostedZoneId"),
			},
			{
				Name:     "iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "latest_restorable_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LatestRestorableTime"),
			},
			{
				Name:     "master_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterUsername"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
			},
			{
				Name:     "percent_progress",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PercentProgress"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "preferred_backup_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredBackupWindow"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "read_replica_identifiers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ReadReplicaIdentifiers"),
			},
			{
				Name:     "reader_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReaderEndpoint"),
			},
			{
				Name:     "replication_source_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationSourceIdentifier"),
			},
			{
				Name:     "serverless_v2_scaling_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ServerlessV2ScalingConfiguration"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "storage_encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("StorageEncrypted"),
			},
			{
				Name:     "vpc_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcSecurityGroups"),
			},
		},
	}
}
