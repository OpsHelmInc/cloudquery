package recipes

import (
	types "github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IdentitystoreResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "groups",
			Struct:     &types.Group{},
			Relations:  []string{"GroupMemberships()"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::IdentityStore::Group")`,
					},
				}...),
		},
		{
			SubService: "users",
			Struct:     &types.User{},
		},
		{
			SubService: "group_memberships",
			Struct:     &types.GroupMembership{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::IdentityStore::GroupMembership")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "identitystore"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
