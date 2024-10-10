package lambda

import (
	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func layerVersionPolicies() *schema.Table {
	tableName := "aws_lambda_layer_version_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersionPolicy.html`,
		Resolver:    fetchLambdaLayerVersionPolicies,
		Transform: transformers.TransformWithStruct(&lambda.GetLayerVersionPolicyOutput{},
			transformers.WithPrimaryKeyComponents("RevisionId"),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "layer_version_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "layer_version",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.ParentColumnResolver("version"),
			},
		},
	}
}
