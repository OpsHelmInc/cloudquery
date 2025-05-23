package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func containerServiceImages() *schema.Table {
	tableName := "aws_lightsail_container_service_images"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_ContainerImage.html`,
		Resolver:    fetchLightsailContainerServiceImages,
		Transform:   transformers.TransformWithStruct(&types.ContainerImage{}, transformers.WithPrimaryKeyComponents("Image")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "container_service_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLightsailContainerServiceImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.ContainerService)
	input := lightsail.GetContainerImagesInput{
		ServiceName: r.ContainerServiceName,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLightsail).Lightsail
	deployments, err := svc.GetContainerImages(ctx, &input, func(options *lightsail.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- deployments.ContainerImages
	return nil
}
