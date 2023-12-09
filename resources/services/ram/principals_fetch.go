package ram

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
)

func fetchRamPrincipals(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	err := fetchRamPrincipalsByOwner(ctx, meta, types.ResourceOwnerSelf, res)
	if err != nil {
		return err
	}
	err = fetchRamPrincipalsByOwner(ctx, meta, types.ResourceOwnerOtherAccounts, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamPrincipalsByOwner(ctx context.Context, meta schema.ClientMeta, shareType types.ResourceOwner, res chan<- interface{}) error {
	input := &ram.ListPrincipalsInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: shareType,
	}
	paginator := ram.NewListPrincipalsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Principals
	}
	return nil
}
