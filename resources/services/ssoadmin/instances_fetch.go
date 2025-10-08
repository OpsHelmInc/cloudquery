package ssoadmin

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
)

func fetchSsoadminInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
