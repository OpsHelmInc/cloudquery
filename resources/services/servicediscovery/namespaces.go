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

func Namespaces() *schema.Table {
	tableName := "aws_servicediscovery_namespaces"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloud-map/latest/api/API_Namespace.html`,
		Resolver:            fetchNamespaces,
		PreResourceResolver: getNamespace,
		Transform:           transformers.TransformWithStruct(&types.Namespace{}),
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
	}
}
func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicediscovery).Servicediscovery
	input := servicediscovery.ListNamespacesInput{MaxResults: aws.Int32(100)}
	paginator := servicediscovery.NewListNamespacesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicediscovery.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Namespaces
	}
	return nil
}

func getNamespace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicediscovery).Servicediscovery
	namespace := resource.Item.(types.NamespaceSummary)

	desc, err := svc.GetNamespace(ctx, &servicediscovery.GetNamespaceInput{Id: namespace.Id}, func(o *servicediscovery.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = desc.Namespace
	return nil
}
