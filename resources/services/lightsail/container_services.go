package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ContainerServices() *schema.Table {
	tableName := "aws_lightsail_container_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerService.html`,
		Resolver:    fetchLightsailContainerServices,
		Transform:   transformers.TransformWithStruct(&types.ContainerService{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			containerServiceDeployments(),
			containerServiceImages(),
		},
	}
}

func fetchLightsailContainerServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetContainerServicesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLightsail).Lightsail
	response, err := svc.GetContainerServices(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.ContainerServices
	return nil
}
