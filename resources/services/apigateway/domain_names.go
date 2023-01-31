// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DomainNames() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_domain_names",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html`,
		Resolver:    fetchApigatewayDomainNames,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Resolver: resolveApigatewayDomainNameArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
			},
			{
				Name:     "certificate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateName"),
			},
			{
				Name:     "certificate_upload_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CertificateUploadDate"),
			},
			{
				Name:     "distribution_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DistributionDomainName"),
			},
			{
				Name:     "distribution_hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DistributionHostedZoneId"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "domain_name_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainNameStatus"),
			},
			{
				Name:     "domain_name_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainNameStatusMessage"),
			},
			{
				Name:     "endpoint_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointConfiguration"),
			},
			{
				Name:     "mutual_tls_authentication",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MutualTlsAuthentication"),
			},
			{
				Name:     "ownership_verification_certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnershipVerificationCertificateArn"),
			},
			{
				Name:     "regional_certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionalCertificateArn"),
			},
			{
				Name:     "regional_certificate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionalCertificateName"),
			},
			{
				Name:     "regional_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionalDomainName"),
			},
			{
				Name:     "regional_hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegionalHostedZoneId"),
			},
			{
				Name:     "security_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SecurityPolicy"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			DomainNameBasePathMappings(),
		},
	}
}
