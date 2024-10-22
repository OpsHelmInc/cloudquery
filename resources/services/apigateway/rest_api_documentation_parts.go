package apigateway

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func restApiDocumentationParts() *schema.Table {
	tableName := "aws_apigateway_rest_api_documentation_parts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPart.html`,
		Resolver:    fetchApigatewayRestApiDocumentationParts,
		Transform:   transformers.TransformWithStruct(&types.DocumentationPart{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveApigatewayRestAPIDocumentationPartArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApigatewayRestApiDocumentationParts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApigateway).Apigateway
	config := apigateway.GetDocumentationPartsInput{RestApiId: r.Id, Limit: aws.Int32(500)}
	// No paginator available
	for {
		response, err := svc.GetDocumentationParts(ctx, &config, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}

func resolveApigatewayRestAPIDocumentationPartArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Item.(types.DocumentationPart)
	rapi := resource.Parent.Item.(types.RestApi)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/restapis/%s/documentation/parts/%s", aws.ToString(rapi.Id), aws.ToString(d.Id)),
	}.String())
}
