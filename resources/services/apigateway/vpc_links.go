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

func VpcLinks() *schema.Table {
	tableName := "aws_apigateway_vpc_links"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html`,
		Resolver:    fetchApigatewayVpcLinks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.VpcLink{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveApigatewayVpcLinkArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApigatewayVpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apigateway.GetVpcLinksInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApigateway).Apigateway
	paginator := apigateway.NewGetVpcLinksPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *apigateway.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}

func resolveApigatewayVpcLinkArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	link := resource.Item.(types.VpcLink)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/vpclinks/%s", aws.ToString(link.Id)),
	}.String())
}
