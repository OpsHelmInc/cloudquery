package networkfirewall

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/networkfirewall/models"
)

func TLSInspectionConfigurations() *schema.Table {
	tableName := "aws_networkfirewall_tls_inspection_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_DescribeTLSInspectionConfiguration.html`,
		Resolver:            fetchTLSInspectionConfigurations,
		PreResourceResolver: getTLSInspectionConfigurations,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.TLSInspectionConfigurationWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TLSInspectionConfigurationArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchTLSInspectionConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListTLSInspectionConfigurationsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkfirewall).Networkfirewall
	p := networkfirewall.NewListTLSInspectionConfigurationsPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- response.TLSInspectionConfigurations
	}
	return nil
}

func getTLSInspectionConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkfirewall).Networkfirewall
	metadata := resource.Item.(types.TLSInspectionConfigurationMetadata)

	tlsInspectionConfigurationDetails, err := svc.DescribeTLSInspectionConfiguration(ctx, &networkfirewall.DescribeTLSInspectionConfigurationInput{
		TLSInspectionConfigurationArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = cl.Region
	})
	if err != nil && !cl.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.TLSInspectionConfigurationWrapper{
		TLSInspectionConfiguration:         tlsInspectionConfigurationDetails.TLSInspectionConfiguration,
		TLSInspectionConfigurationResponse: tlsInspectionConfigurationDetails.TLSInspectionConfigurationResponse,
	}
	return nil
}
