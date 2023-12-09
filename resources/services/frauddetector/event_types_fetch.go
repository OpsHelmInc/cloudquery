package frauddetector

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
)

func fetchFrauddetectorEventTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	paginator := frauddetector.NewGetEventTypesPaginator(meta.(*client.Client).Services().Frauddetector, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.EventTypes
	}
	return nil
}
