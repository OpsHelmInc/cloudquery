package resiliencehub

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func appVersionResourceMappings() *schema.Table {
	tableName := "aws_resiliencehub_app_version_resource_mappings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResourceMapping.html`,
		Resolver:    fetchAppVersionResourceMappings,
		Transform:   transformers.TransformWithStruct(&types.ResourceMapping{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, appVersion,
			{
				Name:                "physical_resource_identifier",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("PhysicalResourceId.Identifier"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAppVersionResourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	p := resiliencehub.NewListAppVersionResourceMappingsPaginator(svc, &resiliencehub.ListAppVersionResourceMappingsInput{
		AppArn:     parent.Parent.Item.(*types.App).AppArn,
		AppVersion: parent.Item.(types.AppVersionSummary).AppVersion,
	})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.ResourceMappings
	}
	return nil
}
