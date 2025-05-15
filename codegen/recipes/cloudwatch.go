package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func CloudwatchResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "alarms",
			Struct:              &ohaws.Alarm{},
			Description:         "https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html",
			SkipFields:          []string{"AlarmArn", "Dimensions", "CompositeAlarm"},
			PreResourceResolver: "getAlarm",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("MetricAlarm.AlarmArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "dimensions",
						Type:     schema.TypeJSON,
						Resolver: `resolveCloudwatchAlarmDimensions`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudwatch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("logs")`
	}
	return resources
}
