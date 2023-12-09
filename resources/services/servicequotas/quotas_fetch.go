package servicequotas

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
)

func fetchServicequotasQuotas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Servicequotas
	service := parent.Item.(types.ServiceInfo)
	config := servicequotas.ListServiceQuotasInput{
		ServiceCode: service.ServiceCode,
		MaxResults:  defaultMaxResults,
	}
	quotasPaginator := servicequotas.NewListServiceQuotasPaginator(svc, &config)
	for quotasPaginator.HasMorePages() {
		output, err := quotasPaginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Quotas
	}
	return nil
}
