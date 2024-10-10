package servicediscovery

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Services() *schema.Table {
	tableName := "aws_servicediscovery_services"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloud-map/latest/api/API_Service.html`,
		Resolver:            fetchServices,
		PreResourceResolver: getService,
		Transform:           transformers.TransformWithStruct(&types.Service{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "servicediscovery"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveServicediscoveryTags("Arn"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			instances(),
		},
	}
}
func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicediscovery).Servicediscovery
	input := servicediscovery.ListServicesInput{MaxResults: aws.Int32(100)}
	paginator := servicediscovery.NewListServicesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicediscovery.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Services
	}
	return nil
}

func getService(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicediscovery).Servicediscovery
	namespace := resource.Item.(types.ServiceSummary)

	desc, err := svc.GetService(ctx, &servicediscovery.GetServiceInput{Id: namespace.Id}, func(o *servicediscovery.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = desc.Service
	return nil
}
