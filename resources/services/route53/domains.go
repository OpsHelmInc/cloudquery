// Code generated by codegen; DO NOT EDIT.

package route53

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "aws_route53_domains",
		Resolver:            fetchRoute53Domains,
		PreResourceResolver: getDomain,
		Multiplex:           client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "domain_name",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveRoute53DomainTags,
				Description: `A list of tags`,
			},
			{
				Name:     "abuse_contact_email",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AbuseContactEmail"),
			},
			{
				Name:     "abuse_contact_phone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AbuseContactPhone"),
			},
			{
				Name:     "admin_contact",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdminContact"),
			},
			{
				Name:     "admin_privacy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AdminPrivacy"),
			},
			{
				Name:     "auto_renew",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoRenew"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "dns_sec",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DnsSec"),
			},
			{
				Name:     "dnssec_keys",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DnssecKeys"),
			},
			{
				Name:     "expiration_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ExpirationDate"),
			},
			{
				Name:     "nameservers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Nameservers"),
			},
			{
				Name:     "registrant_contact",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RegistrantContact"),
			},
			{
				Name:     "registrant_privacy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RegistrantPrivacy"),
			},
			{
				Name:     "registrar_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistrarName"),
			},
			{
				Name:     "registrar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistrarUrl"),
			},
			{
				Name:     "registry_domain_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryDomainId"),
			},
			{
				Name:     "reseller",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Reseller"),
			},
			{
				Name:     "status_list",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("StatusList"),
			},
			{
				Name:     "tech_contact",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TechContact"),
			},
			{
				Name:     "tech_privacy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("TechPrivacy"),
			},
			{
				Name:     "updated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedDate"),
			},
			{
				Name:     "who_is_server",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WhoIsServer"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}
