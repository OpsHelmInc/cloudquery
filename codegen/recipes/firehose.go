package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func FirehoseResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "delivery_streams",
			Struct:              &ohaws.FirehoseStream{},
			Description:         "https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html",
			SkipFields:          []string{"DeliveryStreamARN"},
			PreResourceResolver: "getDeliveryStream",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("DeliveryStreamARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "firehose"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("firehose")`
	}
	return resources
}
