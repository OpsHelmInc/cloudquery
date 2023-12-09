// Code generated by codegen; DO NOT EDIT.

package redshift

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_snapshots",
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchRedshiftSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
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
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveSnapshotARN,
				Description: `ARN of the snapshot.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
				Description: `Tags consisting of a name/value pair for a resource.`,
			},
			{
				Name:     "accounts_with_restore_access",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccountsWithRestoreAccess"),
			},
			{
				Name:     "actual_incremental_backup_size_in_mega_bytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("ActualIncrementalBackupSizeInMegaBytes"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "backup_progress_in_mega_bytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("BackupProgressInMegaBytes"),
			},
			{
				Name:     "cluster_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClusterCreateTime"),
			},
			{
				Name:     "cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterIdentifier"),
			},
			{
				Name:     "cluster_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterVersion"),
			},
			{
				Name:     "current_backup_rate_in_mega_bytes_per_second",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("CurrentBackupRateInMegaBytesPerSecond"),
			},
			{
				Name:     "db_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBName"),
			},
			{
				Name:     "elapsed_time_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ElapsedTimeInSeconds"),
			},
			{
				Name:     "encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Encrypted"),
			},
			{
				Name:     "encrypted_with_hsm",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EncryptedWithHSM"),
			},
			{
				Name:     "engine_full_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineFullVersion"),
			},
			{
				Name:     "enhanced_vpc_routing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnhancedVpcRouting"),
			},
			{
				Name:     "estimated_seconds_to_completion",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EstimatedSecondsToCompletion"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "maintenance_track_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MaintenanceTrackName"),
			},
			{
				Name:     "manual_snapshot_remaining_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ManualSnapshotRemainingDays"),
			},
			{
				Name:     "manual_snapshot_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ManualSnapshotRetentionPeriod"),
			},
			{
				Name:     "master_password_secret_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterPasswordSecretArn"),
			},
			{
				Name:     "master_password_secret_kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterPasswordSecretKmsKeyId"),
			},
			{
				Name:     "master_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterUsername"),
			},
			{
				Name:     "node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NodeType"),
			},
			{
				Name:     "number_of_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfNodes"),
			},
			{
				Name:     "owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerAccount"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "restorable_node_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RestorableNodeTypes"),
			},
			{
				Name:     "snapshot_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SnapshotCreateTime"),
			},
			{
				Name:     "snapshot_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotIdentifier"),
			},
			{
				Name:     "snapshot_retention_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SnapshotRetentionStartTime"),
			},
			{
				Name:     "snapshot_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotType"),
			},
			{
				Name:     "source_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceRegion"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "total_backup_size_in_mega_bytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("TotalBackupSizeInMegaBytes"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
