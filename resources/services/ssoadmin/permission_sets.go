// Code generated by codegen; DO NOT EDIT.

package ssoadmin

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PermissionSets() *schema.Table {
	return &schema.Table{
		Name:                "aws_ssoadmin_permission_sets",
		Description:         `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html`,
		Resolver:            fetchSsoadminPermissionSets,
		PreResourceResolver: getSsoadminPermissionSet,
		Multiplex:           client.ServiceAccountRegionMultiplexer("identitystore"),
		Columns: []schema.Column{
			{
				Name:     "inline_policy",
				Type:     schema.TypeJSON,
				Resolver: getSsoadminPermissionSetInlinePolicy,
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "permission_set_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PermissionSetArn"),
			},
			{
				Name:     "relay_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RelayState"),
			},
			{
				Name:     "session_duration",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SessionDuration"),
			},
		},

		Relations: []*schema.Table{
			AccountAssignments(),
		},
	}
}
