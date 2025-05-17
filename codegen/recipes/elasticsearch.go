package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func ElasticsearchResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "domains",
			Struct:              &ohaws.ElasticsearchDomain{},
			SkipFields:          []string{"DomainId"},
			PreResourceResolver: "getDomain",
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "id",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("DomainId")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elasticsearch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("es")`
	}
	return resources
}
