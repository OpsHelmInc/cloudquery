// Code generated by codegen; DO NOT EDIT.

package accessanalyzer

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AnalyzerFindings() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_findings",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_FindingSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerFindings,
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
				Resolver: resolveFindingArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "analyzer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "analyzed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("AnalyzedAt"),
			},
			{
				Name:     "condition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Condition"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "resource_owner_account",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceOwnerAccount"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "action",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Action"),
			},
			{
				Name:     "error",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Error"),
			},
			{
				Name:     "is_public",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsPublic"),
			},
			{
				Name:     "principal",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Principal"),
			},
			{
				Name:     "resource",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Resource"),
			},
			{
				Name:     "resource_control_policy_restriction",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceControlPolicyRestriction"),
			},
			{
				Name:     "sources",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Sources"),
			},
		},
	}
}
