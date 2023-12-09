package ssm

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func fetchSsmInstancePatches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm
	item := parent.Item.(types.InstanceInformation)

	paginator := ssm.NewDescribeInstancePatchesPaginator(svc, &ssm.DescribeInstancePatchesInput{
		InstanceId: item.InstanceId,
	})
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Patches
	}
	return nil
}
