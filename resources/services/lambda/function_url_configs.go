package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func functionURLConfigs() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_function_url_configs",
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionUrlConfig.html`,
		Resolver:    fetchLambdaFunctionURLConfigs,
		Transform: transformers.TransformWithStruct(&types.FunctionUrlConfig{},
			transformers.WithPrimaryKeyComponents("FunctionArn"), // FunctionArn here can also be a version
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchLambdaFunctionURLConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	config := lambda.ListFunctionUrlConfigsInput{FunctionName: p.Configuration.FunctionArn}
	paginator := lambda.NewListFunctionUrlConfigsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.FunctionUrlConfigs
	}
	return nil
}
