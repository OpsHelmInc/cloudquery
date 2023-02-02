// Code generated by codegen; DO NOT EDIT.

package identitystore

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupMemberships() *schema.Table {
	return &schema.Table{
		Name:      "aws_identitystore_group_memberships",
		Resolver:  fetchIdentitystoreGroupMemberships,
		Multiplex: client.ServiceAccountRegionMultiplexer("identitystore"),
		Columns: []schema.Column{
			{
				Name:     "identity_store_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityStoreId"),
			},
			{
				Name:     "group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupId"),
			},
			{
				Name:     "membership_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MembershipId"),
			},
		},
	}
}
