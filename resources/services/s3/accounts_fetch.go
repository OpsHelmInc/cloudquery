package s3

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/OpsHelmInc/ohaws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/pkg/errors"
)

func fetchS3Accounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().S3control
	var accountConfig s3control.GetPublicAccessBlockInput
	accountConfig.AccountId = aws.String(c.AccountID)
	resp, err := svc.GetPublicAccessBlock(ctx, &accountConfig)

	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		var nspabc *s3controlTypes.NoSuchPublicAccessBlockConfiguration
		if !errors.As(err, &nspabc) {
			return err
		}
		res <- ohaws.AccountPublicAccessBlockConfigurationWrapper{ConfigExists: false}
	} else {
		res <- ohaws.AccountPublicAccessBlockConfigurationWrapper{PublicAccessBlockConfiguration: *resp.PublicAccessBlockConfiguration, ConfigExists: true}
	}

	return nil
}
