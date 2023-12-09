// Code generated by codegen; DO NOT EDIT.

package servicecatalog

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "aws_servicecatalog_products",
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_ProductViewDetail.html`,
		Resolver:    fetchServicecatalogProducts,
		Multiplex:   client.ServiceAccountRegionMultiplexer("servicecatalog"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProductARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveProductTags,
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "product_view_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProductViewSummary"),
			},
			{
				Name:     "source_connection",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceConnection"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
