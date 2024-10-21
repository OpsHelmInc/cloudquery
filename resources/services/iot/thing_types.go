package iot

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ThingTypes() *schema.Table {
	tableName := "aws_iot_thing_types"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_ThingTypeDefinition.html`,
		Resolver:    fetchIotThingTypes,
		Transform:   transformers.TransformWithStruct(&types.ThingTypeDefinition{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: ResolveIotThingTypeTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ThingTypeArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIotThingTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingTypesInput{
		MaxResults: aws.Int32(250),
	}
	cl := meta.(*client.Client)

	svc := cl.Services(client.AWSServiceIot).Iot
	paginator := iot.NewListThingTypesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- page.ThingTypes
	}
	return nil
}

func ResolveIotThingTypeTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(types.ThingTypeDefinition)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIot).Iot
	return resolveIotTags(ctx, meta, svc, resource, c, i.ThingTypeArn)
}
