// Code generated by codegen; DO NOT EDIT.

package kms

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:                "aws_kms_keys",
		Description:         `https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html`,
		Resolver:            fetchKmsKeys,
		PreResourceResolver: getKey,
		Multiplex:           client.ServiceAccountRegionMultiplexer("kms"),
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
				Name:     "rotation_enabled",
				Type:     schema.TypeBool,
				Resolver: resolveKeysRotationEnabled,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKeysTags,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:          "replica_keys",
				Type:          schema.TypeJSON,
				Resolver:      resolveKeysReplicaKeys,
				IgnoreInTests: true,
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::KMS::Key"),
			},
			{
				Name:     "key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyId"),
			},
			{
				Name:     "aws_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AWSAccountId"),
			},
			{
				Name:     "cloud_hsm_cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudHsmClusterId"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "custom_key_store_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomKeyStoreId"),
			},
			{
				Name:     "customer_master_key_spec",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerMasterKeySpec"),
			},
			{
				Name:     "deletion_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletionDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "encryption_algorithms",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EncryptionAlgorithms"),
			},
			{
				Name:     "expiration_model",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpirationModel"),
			},
			{
				Name:     "key_manager",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyManager"),
			},
			{
				Name:     "key_spec",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeySpec"),
			},
			{
				Name:     "key_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyState"),
			},
			{
				Name:     "key_usage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyUsage"),
			},
			{
				Name:     "mac_algorithms",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("MacAlgorithms"),
			},
			{
				Name:     "multi_region",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MultiRegion"),
			},
			{
				Name:     "multi_region_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MultiRegionConfiguration"),
			},
			{
				Name:     "origin",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Origin"),
			},
			{
				Name:     "pending_deletion_window_in_days",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PendingDeletionWindowInDays"),
			},
			{
				Name:     "signing_algorithms",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SigningAlgorithms"),
			},
			{
				Name:     "valid_to",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ValidTo"),
			},
			{
				Name:     "xks_key_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("XksKeyConfiguration"),
			},
		},

		Relations: []*schema.Table{
			KeyGrants(),
		},
	}
}
