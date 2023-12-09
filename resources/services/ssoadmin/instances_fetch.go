package ssoadmin

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
)

func fetchSsoadminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Ssoadmin
	config := ssoadmin.ListInstancesInput{}
	response, err := svc.ListInstances(ctx, &config)
	if err != nil {
		return err
	}
	for _, i := range response.Instances {
		res <- i
	}
	return nil
}
