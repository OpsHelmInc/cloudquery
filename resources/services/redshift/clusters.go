// Code generated by codegen; DO NOT EDIT.

package redshift

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_clusters",
		Description: `https://docs.aws.amazon.com/redshift/latest/APIReference/API_Cluster.html`,
		Resolver:    fetchRedshiftClusters,
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
				Resolver:    resolveClusterArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "logging_status",
				Type:        schema.TypeJSON,
				Resolver:    resolveRedshiftClusterLoggingStatus,
				Description: `Describes the status of logging for a cluster.`,
			},
			{
				Name:     "allow_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AllowVersionUpgrade"),
			},
			{
				Name:     "aqua_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AquaConfiguration"),
			},
			{
				Name:     "automated_snapshot_retention_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AutomatedSnapshotRetentionPeriod"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "availability_zone_relocation_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneRelocationStatus"),
			},
			{
				Name:     "cluster_availability_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterAvailabilityStatus"),
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
				Name:     "cluster_namespace_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterNamespaceArn"),
			},
			{
				Name:     "cluster_nodes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterNodes"),
			},
			{
				Name:     "cluster_public_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterPublicKey"),
			},
			{
				Name:     "cluster_revision_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterRevisionNumber"),
			},
			{
				Name:     "cluster_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterSecurityGroups"),
			},
			{
				Name:     "cluster_snapshot_copy_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ClusterSnapshotCopyStatus"),
			},
			{
				Name:     "cluster_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterStatus"),
			},
			{
				Name:     "cluster_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterSubnetGroupName"),
			},
			{
				Name:     "cluster_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterVersion"),
			},
			{
				Name:     "custom_domain_certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomDomainCertificateArn"),
			},
			{
				Name:     "custom_domain_certificate_expiry_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CustomDomainCertificateExpiryDate"),
			},
			{
				Name:     "custom_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomDomainName"),
			},
			{
				Name:     "db_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DBName"),
			},
			{
				Name:     "data_transfer_progress",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DataTransferProgress"),
			},
			{
				Name:     "default_iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultIamRoleArn"),
			},
			{
				Name:     "deferred_maintenance_windows",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DeferredMaintenanceWindows"),
			},
			{
				Name:     "elastic_ip_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ElasticIpStatus"),
			},
			{
				Name:     "elastic_resize_number_of_node_options",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ElasticResizeNumberOfNodeOptions"),
			},
			{
				Name:     "encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Encrypted"),
			},
			{
				Name:     "endpoint",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Endpoint"),
			},
			{
				Name:     "enhanced_vpc_routing",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnhancedVpcRouting"),
			},
			{
				Name:     "expected_next_snapshot_schedule_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExpectedNextSnapshotScheduleTime"),
			},
			{
				Name:     "expected_next_snapshot_schedule_time_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpectedNextSnapshotScheduleTimeStatus"),
			},
			{
				Name:     "hsm_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HsmStatus"),
			},
			{
				Name:     "iam_roles",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IamRoles"),
			},
			{
				Name:     "ip_address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddressType"),
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
				Name:     "modify_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModifyStatus"),
			},
			{
				Name:     "multi_az",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MultiAZ"),
			},
			{
				Name:     "multi_az_secondary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MultiAZSecondary"),
			},
			{
				Name:     "next_maintenance_window_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NextMaintenanceWindowStartTime"),
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
				Name:     "pending_actions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("PendingActions"),
			},
			{
				Name:     "pending_modified_values",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PendingModifiedValues"),
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
				Name:     "reserved_node_exchange_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ReservedNodeExchangeStatus"),
			},
			{
				Name:     "resize_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResizeInfo"),
			},
			{
				Name:     "restore_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RestoreStatus"),
			},
			{
				Name:     "snapshot_schedule_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotScheduleIdentifier"),
			},
			{
				Name:     "snapshot_schedule_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotScheduleState"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "total_storage_capacity_in_mega_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TotalStorageCapacityInMegaBytes"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
			{
				Name:     "vpc_security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcSecurityGroups"),
			},
		},

		Relations: []*schema.Table{
			Snapshots(),
			ClusterParameterGroups(),
		},
	}
}
