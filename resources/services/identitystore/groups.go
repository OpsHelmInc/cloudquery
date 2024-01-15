// Code generated by codegen; DO NOT EDIT.

package identitystore

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "aws_identitystore_groups",
		Resolver:  fetchIdentitystoreGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("identitystore"),
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
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::IdentityStore::Group"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupId"),
			},
			{
				Name:     "identity_store_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityStoreId"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "external_ids",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExternalIds"),
			},
		},

		Relations: []*schema.Table{
			GroupMemberships(),
		},
	}
}
