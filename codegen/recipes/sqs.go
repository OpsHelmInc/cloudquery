package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/resources/services/sqs/models"
)

func SQSResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "queues",
			Struct:              &models.Queue{},
			SkipFields:          []string{"Arn", "Policy", "RedriveAllowPolicy", "RedrivePolicy"},
			PreResourceResolver: "getQueue",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveSqsQueueTags`,
					},
					{
						Name:     "policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("Policy")`,
					},
					{
						Name:     "redrive_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("RedrivePolicy")`,
					},
					{
						Name:     "redrive_allow_policy",
						Type:     schema.TypeJSON,
						Resolver: `schema.PathResolver("RedriveAllowPolicy")`,
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::SQS::Queue")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "sqs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("sqs")`
	}
	return resources
}
