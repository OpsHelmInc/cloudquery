// Code generated by codegen; DO NOT EDIT.

package rds

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func ClusterSnapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_rds_cluster_snapshots",
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterSnapshot.html`,
		Resolver:    fetchRdsClusterSnapshots,
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
				Resolver: schema.PathResolver("DBClusterSnapshotArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRDSClusterSnapshotTags,
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: resolveRDSClusterSnapshotAttributes,
			},
			{
				Name:     "allocated_storage",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AllocatedStorage"),
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "cluster_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ClusterCreateTime"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_cluster_snapshot_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterSnapshotIdentifier"),
			},
			{
				Name:     "db_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBSystemId"),
			},
			{
				Name:     "db_cluster_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DbClusterResourceId"),
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
				Name:     "license_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LicenseModel"),
			},
			{
				Name:     "master_username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MasterUsername"),
			},
			{
				Name:     "percent_progress",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PercentProgress"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "snapshot_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SnapshotCreateTime"),
			},
			{
				Name:     "snapshot_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotType"),
			},
			{
				Name:     "source_db_cluster_snapshot_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SourceDBClusterSnapshotArn"),
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
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
