package lambda

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/resources/services/lambda/models"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	runtimes := make([]models.RuntimeWrapper, len(types.RuntimeProvidedal2.Values()))
	for i, runtime := range types.RuntimeProvidedal2.Values() {
		runtimes[i] = models.RuntimeWrapper{
			Name: string(runtime),
		}
	}
	res <- runtimes
	return nil
}
