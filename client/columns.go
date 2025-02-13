package client

import (
	"github.com/apache/arrow/go/v16/arrow"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

func DefaultAccountIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "account_id",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSAccount,
		PrimaryKeyComponent: pk,
	}
}

func RequestAccountIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "request_account_id",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSAccount,
		PrimaryKeyComponent: pk,
	}
}

func DefaultRegionColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "region",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSRegion,
		PrimaryKeyComponent: pk,
	}
}

func RequestRegionColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "request_region",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveAWSRegion,
		PrimaryKeyComponent: pk,
	}
}

func LanguageCodeColumn(pk bool) schema.Column {
	return schema.Column{
		Name:                "language_code",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveLanguageCode,
		PrimaryKeyComponent: pk,
	}
}

func OhResourceTypeColumn() schema.Column {
	return schema.Column{
		Name:                "oh_resource_type",
		Type:                arrow.BinaryTypes.String,
		Resolver:            ResolveOHResourceType,
		PrimaryKeyComponent: false,
	}
}
