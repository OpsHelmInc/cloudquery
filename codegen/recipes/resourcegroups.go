package recipes

import (
	"reflect"
	"strings"

	"github.com/OpsHelmInc/cloudquery/resources/services/resourcegroups/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceGroupsResources() []*Resource {
	resources := []*Resource{

		{
			SubService:          "resource_groups",
			Struct:              &models.ResourceGroupWrapper{},
			SkipFields:          []string{},
			PreResourceResolver: "getResourceGroup",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("GroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveResourcegroupsResourceGroupTags`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ResourceGroups::Group")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "resourcegroups"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("resource-groups")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
