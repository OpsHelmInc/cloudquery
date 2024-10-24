package lambda

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func functionVersions() *schema.Table {
	tableName := "aws_lambda_function_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html`,
		Resolver:    fetchLambdaFunctionVersions,
		Transform:   transformers.TransformWithStruct(&types.FunctionConfiguration{}, transformers.WithPrimaryKeyComponents("Version")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "function_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLambdaFunctionVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	config := lambda.ListVersionsByFunctionInput{FunctionName: p.Configuration.FunctionArn}
	paginator := lambda.NewListVersionsByFunctionPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if meta.(*client.Client).IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.Versions
	}
	return nil
}
