package servicequotas

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
)

func fetchServicequotasServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := servicequotas.ListServicesInput{
		MaxResults: defaultMaxResults,
	}

	svc := meta.(*client.Client).Services().Servicequotas
	servicePaginator := servicequotas.NewListServicesPaginator(svc, &config)
	for servicePaginator.HasMorePages() {
		output, err := servicePaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Services
	}
	return nil
}
