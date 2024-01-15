package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SchedulerResources() []*Resource {
	mx := `client.ServiceAccountRegionMultiplexer("scheduler")`
	resources := []*Resource{
		{
			SubService: "schedule_groups",
			Struct:     new(types.ScheduleGroupSummary),
			Multiplex:  mx,
			PKColumns:  []string{"arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{

					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveSchedulerScheduleTags()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::Scheduler::ScheduleGroup")`,
					},
				}...),
		},
		{
			SubService: "schedules",
			Struct:     new(scheduler.GetScheduleOutput),
			Multiplex:  mx,
			PKColumns:  []string{"arn"},
			SkipFields: []string{"ResultMetadata"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{

					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveSchedulerScheduleTags()`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::Scheduler::Schedule")`,
					},
				}...),
			PreResourceResolver: "getSchedule",
		},
	}
	for _, r := range resources {
		r.Service = "scheduler"
		r.Description = "https://docs.aws.amazon.com/scheduler/latest/APIReference/API_" + r.StructName() + ".html"
	}
	return resources
}
