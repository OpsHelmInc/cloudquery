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
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveIdentityStoreGroupArn`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     ohResourceTypeColumn,
					Type:     schema.TypeString,
					Resolver: `client.StaticValueResolver("AWS::IdentityStore::Group")`,
				},
			},
		},
		{
			SubService: "users",
			Struct:     &types.User{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `resolveIdentityStoreUserArn`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     ohResourceTypeColumn,
					Type:     schema.TypeString,
					Resolver: `client.StaticValueResolver("AWS::IdentityStore::User")`,
				},
			},
		},
		{
			SubService: "group_memberships",
			Struct:     &types.GroupMembership{},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "identitystore"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("identitystore")`
	}
	return resources
}
