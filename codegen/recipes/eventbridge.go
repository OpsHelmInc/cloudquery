package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func EventbridgeResources() []*Resource {
	regionalResources := []*Resource{
		{
			SubService:            "event_buses",
			Struct:                &ohaws.EventBus{},
			PreResourceResolver:   "getEventbridgeEventBus",
			Description:           "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html",
			SkipFields:            []string{"Arn"},
			UnwrapEmbeddedStructs: true,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "arn",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"EventBusRules()",
			},
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:            "event_bus_rules",
			Struct:                &ohaws.EventBridgeRule{},
			PreResourceResolver:   "getEventBusRule",
			Description:           "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Rule.html",
			SkipFields:            []string{},
			UnwrapEmbeddedStructs: true,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "event_bus_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:            "api_destinations",
			Struct:                &ohaws.EventBridgeAPIDestination{},
			PreResourceResolver:   "getEventBridgeAPIDestination",
			Description:           "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html",
			SkipFields:            []string{"ApiDestinationArn"},
			UnwrapEmbeddedStructs: true,
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("ApiDestinationArn")`,
				},
			}...),
			ShouldGenerateResolverAndMockTest: false,
		},
		{
			SubService:  "archives",
			Struct:      &types.Archive{},
			Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Archive.html",
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{{
				Name:     "arn",
				Type:     schema.TypeString,
				Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				Resolver: `resolveArchiveArn`,
			}}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "connections",
			Struct:      &types.Connection{},
			Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Connection.html",
			SkipFields:  []string{"ConnectionArn"},
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("ConnectionArn")`,
				},
			}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "event_sources",
			Struct:      &types.EventSource{},
			Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventSource.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Arn")`,
				},
			}...),
			ShouldGenerateResolverAndMockTest: true,
		},
		{
			SubService:  "replays",
			Struct:      &types.Replay{},
			Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Replay.html",
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `resolveReplayArn`,
				},
			}...),
			ShouldGenerateResolverAndMockTest: true,
		},
	}

	globalResources := []*Resource{
		{
			SubService:  "endpoints",
			Struct:      &types.Endpoint{},
			Description: "https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(defaultAccountColumns, []codegen.ColumnDefinition{
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					Resolver: `schema.PathResolver("Arn")`,
				},
			}...),
			ShouldGenerateResolverAndMockTest: true,
		},
	}

	// set default values
	for _, r := range regionalResources {
		r.Service = "eventbridge"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("events")`

		// Parameters for autogenerating the resolver and mock-test.
		/// Only used when `ShouldGenerateResolverAndMockTest = true`
		r.ResolverAndMockTestTemplate = "list_resources_1"
		r.Client = &eventbridge.Client{}
	}

	for _, r := range globalResources {
		r.Service = "eventbridge"
		r.Multiplex = `client.AccountMultiplex`

		// Parameters for autogenerating the resolver and mock-test.
		/// Only used when `ShouldGenerateResolverAndMockTest = true`
		r.ResolverAndMockTestTemplate = "list_resources_1"
		r.Client = &eventbridge.Client{}
	}

	resources := append(regionalResources, globalResources...)

	return resources
}
