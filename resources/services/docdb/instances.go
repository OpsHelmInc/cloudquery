// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_instances",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBInstance.html`,
		Resolver:    fetchDocdbInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveDBInstanceTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "backup_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BackupRetentionPeriod"),
			},
			{
				Name:     "ca_certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CACertificateIdentifier"),
			},
			{
				Name:     "certificate_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CertificateDetails"),
			},
			{
				Name:     "copy_tags_to_snapshot",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CopyTagsToSnapshot"),
			},
			{
				Name:     "db_cluster_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBClusterIdentifier"),
			},
			{
				Name:     "db_instance_class",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceClass"),
			},
			{
				Name:     "db_instance_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceIdentifier"),
			},
			{
				Name:     "db_instance_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBInstanceStatus"),
			},
			{
				Name:     "db_subnet_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DBSubnetGroup"),
			},
			{
				Name:     "dbi_resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DbiResourceId"),
			},
			{
				Name:     "enabled_cloudwatch_logs_exports",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EnabledCloudwatchLogsExports"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeJSON,
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
				Name:     "instance_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("InstanceCreateTime"),
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
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
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
				Name:     "promotion_tier",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PromotionTier"),
			},
			{
				Name:     "publicly_accessible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PubliclyAccessible"),
			},
			{
				Name:     "status_infos",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatusInfos"),
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
