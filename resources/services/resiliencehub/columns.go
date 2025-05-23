package resiliencehub

import (
	"github.com/apache/arrow/go/v16/arrow"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
)

var (
	appARN = schema.Column{
		Name:                "app_arn",
		Type:                arrow.BinaryTypes.String,
		Resolver:            schema.ParentColumnResolver("app_arn"),
		PrimaryKeyComponent: true,
	}
	appARNTop = schema.Column{
		Name:                "app_arn",
		Type:                arrow.BinaryTypes.String,
		Resolver:            schema.ParentColumnResolver("arn"),
		PrimaryKeyComponent: true,
	}
	assessmentARN = schema.Column{
		Name:                "assessment_arn",
		Type:                arrow.BinaryTypes.String,
		Resolver:            schema.ParentColumnResolver("arn"),
		PrimaryKeyComponent: true,
	}
	appVersion = schema.Column{
		Name:                "app_version",
		Type:                arrow.BinaryTypes.String,
		Resolver:            schema.ParentColumnResolver("app_version"),
		PrimaryKeyComponent: true,
	}
)

func arnColumn(path string) schema.Column {
	return schema.Column{
		Name:                "arn",
		Type:                arrow.BinaryTypes.String,
		Resolver:            schema.PathResolver(path),
		PrimaryKeyComponent: true,
	}
}
