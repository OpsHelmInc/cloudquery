package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func transitGatewayRegistration() *schema.Table {
	tableName := "aws_networkmanager_transit_gateway_registrations"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/networkmanager/latest/APIReference/API_TransitGatewayRegistration.html
The  'request_region' column is added to show region of where the request was made from.`,
		Resolver:  fetchTransitGatewayRegistration,
		Transform: transformers.TransformWithStruct(&types.TransitGatewayRegistration{}, transformers.WithPrimaryKeyComponents("GlobalNetworkId", "TransitGatewayArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.RequestRegionColumn(true),
		},
		Relations: schema.Tables{},
	}
}

func fetchTransitGatewayRegistration(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkmanager).Networkmanager
	globalNetwork := parent.Item.(types.GlobalNetwork)
	input := &networkmanager.GetTransitGatewayRegistrationsInput{
		GlobalNetworkId: globalNetwork.GlobalNetworkId,
	}
	paginator := networkmanager.NewGetTransitGatewayRegistrationsPaginator(svc, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *networkmanager.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TransitGatewayRegistrations
	}
	return nil
}
