package lambda

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func layerVersions() *schema.Table {
	tableName := "aws_lambda_layer_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_LayerVersionsListItem.html`,
		Resolver:    fetchLambdaLayerVersions,
		Transform:   transformers.TransformWithStruct(&types.LayerVersionsListItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("LayerVersionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "layer_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			layerVersionPolicies(),
		},
	}
}
