package apprunner

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
)

func fetchApprunnerOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := apprunner.NewListOperationsPaginator(meta.(*client.Client).Services().Apprunner,
		&apprunner.ListOperationsInput{ServiceArn: parent.Item.(*types.Service).ServiceArn})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.OperationSummaryList
	}
	return nil
}
