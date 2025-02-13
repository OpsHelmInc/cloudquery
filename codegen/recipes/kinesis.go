package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func KinesisResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "streams",
			Struct:              &ohaws.KinesisStream{},
			Description:         "https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html",
			SkipFields:          []string{"StreamARN"},
			PreResourceResolver: "getStream",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("StreamARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "kinesis"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("kinesis")`
	}
	return resources
}
