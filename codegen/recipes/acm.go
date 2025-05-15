package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func ACMResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "certificates",
			Struct:              &ohaws.ACMCertificate{},
			Description:         "https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html",
			SkipFields:          []string{"CertificateArn"},
			PreResourceResolver: "getCertificate",
			Multiplex:           `client.ServiceAccountRegionMultiplexer("acm")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("CertificateArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "acm"
	}
	return resources
}
