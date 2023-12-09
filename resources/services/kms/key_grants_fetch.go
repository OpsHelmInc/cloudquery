package kms

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

func fetchKmsKeyGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	k := parent.Item.(*types.KeyMetadata)
	config := kms.ListGrantsInput{
		KeyId: k.Arn,
		Limit: aws.Int32(100),
	}

	c := meta.(*client.Client)
	svc := c.Services().Kms
	p := kms.NewListGrantsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Grants
	}
	return nil
}
