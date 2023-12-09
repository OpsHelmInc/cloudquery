// Code generated by codegen; DO NOT EDIT.

package glue

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func DatabaseTables() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_database_tables",
		Resolver:  fetchGlueDatabaseTables,
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:     "database_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "catalog_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CatalogId"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "created_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedBy"),
			},
			{
				Name:     "database_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DatabaseName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "federated_table",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FederatedTable"),
			},
			{
				Name:     "is_registered_with_lake_formation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsRegisteredWithLakeFormation"),
			},
			{
				Name:     "last_access_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAccessTime"),
			},
			{
				Name:     "last_analyzed_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastAnalyzedTime"),
			},
			{
				Name:     "owner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Owner"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "partition_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PartitionKeys"),
			},
			{
				Name:     "retention",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Retention"),
			},
			{
				Name:     "storage_descriptor",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StorageDescriptor"),
			},
			{
				Name:     "table_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TableType"),
			},
			{
				Name:     "target_table",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetTable"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdateTime"),
			},
			{
				Name:     "version_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VersionId"),
			},
			{
				Name:     "view_expanded_text",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewExpandedText"),
			},
			{
				Name:     "view_original_text",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewOriginalText"),
			},
		},

		Relations: []*schema.Table{
			DatabaseTableIndexes(),
		},
	}
}
