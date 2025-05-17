package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/ohaws"
)

func ELBv2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "load_balancers",
			Struct:              &ohaws.LoadBalancerV2{},
			PreResourceResolver: "getLoadBalancer",
			Description:         "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html",
			SkipFields: []string{
				"LoadBalancerArn",
				"DeletionProtectionEnabled",
				"CrossZoneLoadBalancingEnabled",
				"AccessLogsEnabled",
				"AccessLogsS3Bucket",
				"AccessLogsS3Prefix",
				"IPv6DenyAllIGWTraffic",
				"IdleTimeoutSeconds",
				"ClientKeepAliveSeconds",
				"ConnectionLogsEnabled",
				"ConnectionLogsS3Bucket",
				"ConnectionLogsS3Prefix",
				"DesyncMitigationMode",
				"DropInvalidHeaderFields",
				"PreserveHostHeader",
				"XAmznTLSVersionAndCipherSuiteEnabled",
				"XffClientPortEnabled",
				"XffHeaderProcessingMode",
				"HTTP2Enabled",
				"WafFailOpen",
				"ClientRoutingPolicy",
			},
			UnwrapEmbeddedStructs: false,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:          "web_acl_arn",
						Type:          schema.TypeString,
						Resolver:      `resolveElbv2loadBalancerWebACLArn`,
						IgnoreInTests: true,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LoadBalancerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ElasticLoadBalancingV2::LoadBalancer")`,
					},
				}...),
			Relations: []string{
				"Listeners()",
			},
		},
		{
			SubService:            "listeners",
			Struct:                &ohaws.LoadBalancerV2Listener{},
			PreResourceResolver:   "getLoadBalancerListener",
			Description:           "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Listener.html",
			SkipFields:            []string{"ListenerArn"},
			UnwrapEmbeddedStructs: false,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ListenerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ElasticLoadBalancingV2::Listener")`,
					},
				}...),
			Relations: []string{
				"ListenerCertificates()",
			},
		},
		{
			SubService:  "listener_certificates",
			Struct:      &types.Certificate{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_Certificate.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "listener_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:          "target_groups",
			Struct:              &ohaws.TargetGroup{},
			PreResourceResolver: "getTargetGroup",
			Description:         "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html",
			SkipFields: []string{
				"TargetGroupArn",
				"DeregistrationDelaySeconds",
				"StickinessEnabled",
				"StickinessType",
				"CrossZoneLoadBalancingEnabled",
				"DNSFailOverMinimumHealthyTargetCount",
				"DNSFailOverMinimumHealthyTargetPercentage",
				"RoutingFailoverMinimumHealthyTargetCount",
				"RoutingFailoverMinimumHealthyTargetPercentage",
				"LoadBalancingAlgorithm",
				"AnomalyMitigation",
				"SlowStartSeconds",
				"AppStickinessCookieName",
				"AppStickinessDurationSeconds",
				"LoadBalancerStickinessDurationSeconds",
				"LambdaMultiValueHeadersEnabled",
				"DeregistrationConnectionTerminationEnabled",
				"PreserveClientIP",
				"ProxyProtocolV2Enabled",
				"UnhealthyConnectionTerminatationEnabled",
				"UnhealthyConnectionDrainingIntervalSeconds",
				"TargetFailOverOnDeregistration",
				"TargetFailOverOnUnhealthy",
			},
			UnwrapEmbeddedStructs: false,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TargetGroupArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     ohResourceTypeColumn,
						Type:     schema.TypeString,
						Resolver: `client.StaticValueResolver("AWS::ElasticLoadBalancingV2::TargetGroup")`,
					},
				}...),
			Relations: []string{
				"TargetGroupTargetHealthDescriptions()",
			},
		},
		{
			SubService:  "target_group_target_health_descriptions",
			Struct:      &types.TargetHealthDescription{},
			Description: "https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetHealthDescription.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "target_group_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "elbv2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("elasticloadbalancing")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
