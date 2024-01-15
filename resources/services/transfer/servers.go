// Code generated by codegen; DO NOT EDIT.

package transfer

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:                "aws_transfer_servers",
		Description:         `https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html`,
		Resolver:            fetchTransferServers,
		PreResourceResolver: getServer,
		Multiplex:           client.ServiceAccountRegionMultiplexer("transfer"),
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
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveServersTags,
				Description: `Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described`,
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::Transfer::Server"),
			},
			{
				Name:     "certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Certificate"),
			},
			{
				Name:     "domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Domain"),
			},
			{
				Name:     "endpoint_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointDetails"),
			},
			{
				Name:     "endpoint_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointType"),
			},
			{
				Name:     "host_key_fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostKeyFingerprint"),
			},
			{
				Name:     "identity_provider_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProviderDetails"),
			},
			{
				Name:     "identity_provider_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IdentityProviderType"),
			},
			{
				Name:     "logging_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoggingRole"),
			},
			{
				Name:     "post_authentication_login_banner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PostAuthenticationLoginBanner"),
			},
			{
				Name:     "pre_authentication_login_banner",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreAuthenticationLoginBanner"),
			},
			{
				Name:     "protocol_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProtocolDetails"),
			},
			{
				Name:     "protocols",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Protocols"),
			},
			{
				Name:     "s3_storage_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("S3StorageOptions"),
			},
			{
				Name:     "security_policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityPolicyName"),
			},
			{
				Name:     "server_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerId"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "structured_log_destinations",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("StructuredLogDestinations"),
			},
			{
				Name:     "user_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UserCount"),
			},
			{
				Name:     "workflow_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WorkflowDetails"),
			},
		},
	}
}
