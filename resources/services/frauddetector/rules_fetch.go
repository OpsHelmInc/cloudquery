package frauddetector

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
)

func fetchFrauddetectorRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := frauddetector.NewGetRulesPaginator(meta.(*client.Client).Services().Frauddetector,
		&frauddetector.GetRulesInput{DetectorId: parent.Item.(types.Detector).DetectorId})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.RuleDetails
	}
	return nil
}
