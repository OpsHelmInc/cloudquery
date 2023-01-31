// Code generated by codegen; DO NOT EDIT.

package resourcegroups

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceGroups() *schema.Table {
	return &schema.Table{
		Name:                "aws_resourcegroups_resource_groups",
		Resolver:            fetchResourcegroupsResourceGroups,
		PreResourceResolver: getResourceGroup,
		Multiplex:           client.ServiceAccountRegionMultiplexer("resource-groups"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcegroupsResourceGroupTags,
			},
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GroupArn"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "query",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Query"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}
