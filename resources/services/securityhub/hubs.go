package securityhub

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Hubs() *schema.Table {
	tableName := "aws_securityhub_hubs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html`,
		Resolver:    fetchHubs,
		Transform: transformers.TransformWithStruct(&securityhub.DescribeHubOutput{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: fetchHubTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("HubArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchHubs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSecurityhub).Securityhub
	hub, err := svc.DescribeHub(ctx, nil, func(o *securityhub.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- hub
	return nil
}

func fetchHubTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSecurityhub).Securityhub
	config := &securityhub.ListTagsForResourceInput{ResourceArn: resource.Item.(*securityhub.DescribeHubOutput).HubArn}
	tags, err := svc.ListTagsForResource(ctx, config, func(o *securityhub.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, tags.Tags)
}
