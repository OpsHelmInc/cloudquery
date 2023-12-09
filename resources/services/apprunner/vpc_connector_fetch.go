package apprunner

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
)

func fetchApprunnerVpcConnectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apprunner.ListVpcConnectorsInput
	svc := meta.(*client.Client).Services().Apprunner
	paginator := apprunner.NewListVpcConnectorsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.VpcConnectors
	}
	return nil
}
