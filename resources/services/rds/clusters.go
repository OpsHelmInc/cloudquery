// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_clusters",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBCluster.html`,
		Resolver:    fetchRdsClusters,
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
				Resolver: resolveRdsClusterTags,
			},
			{
				Name:     "activity_stream_kinesis_stream_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamKinesisStreamName"),
			},
			{
				Name:     "activity_stream_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamKmsKeyId"),
			},
			{
				Name:     "activity_stream_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamMode"),
			},
			{
				Name:     "activity_stream_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityStreamStatus"),
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
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
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
				Name:     "backtrack_consumed_change_records",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BacktrackConsumedChangeRecords"),
			},
			{
				Name:     "backtrack_window",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BacktrackWindow"),
			},
			{
				Name:     "backup_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BackupRetentionPeriod"),
			},
			{
				Name:     "capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Capacity"),
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
				Name:     "custom_endpoints",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CustomEndpoints"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterInstanceClass"),
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
				Name:     "db_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSystemId"),
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
				Name:     "domain_memberships",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainMemberships"),
			},
			{
				Name:     "earliest_backtrack_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EarliestBacktrackTime"),
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
				Name:     "engine_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineMode"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "global_write_forwarding_requested",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("GlobalWriteForwardingRequested"),
			},
			{
				Name:     "global_write_forwarding_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GlobalWriteForwardingStatus"),
			},
			{
				Name:     "hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostedZoneId"),
			},
			{
				Name:     "http_endpoint_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HttpEndpointEnabled"),
			},
			{
				Name:     "iam_database_authentication_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IAMDatabaseAuthenticationEnabled"),
			},
			{
				Name:     "io_optimized_next_allowed_modification_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("IOOptimizedNextAllowedModificationTime"),
			},
			{
				Name:     "iops",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Iops"),
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
				Name:     "master_user_secret",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MasterUserSecret"),
			},
			{
				Name:     "master_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterUsername"),
			},
			{
				Name:     "monitoring_interval",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MonitoringInterval"),
			},
			{
				Name:     "monitoring_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonitoringRoleArn"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "network_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NetworkType"),
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
				Name:     "performance_insights_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PerformanceInsightsEnabled"),
			},
			{
				Name:     "performance_insights_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PerformanceInsightsKMSKeyId"),
			},
			{
				Name:     "performance_insights_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PerformanceInsightsRetentionPeriod"),
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
				Name:     "publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PubliclyAccessible"),
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
				Name:     "scaling_configuration_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ScalingConfigurationInfo"),
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
				Name:     "storage_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StorageType"),
			},
			{
				Name:     "vpc_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcSecurityGroups"),
			},
		},
	}
}
