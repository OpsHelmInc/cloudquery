package frauddetector

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFrauddetectorModels(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	paginator := frauddetector.NewGetModelsPaginator(meta.(*client.Client).Services().Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Models
	}
	return nil
}
