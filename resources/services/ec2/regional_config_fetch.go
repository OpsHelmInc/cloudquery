package ec2

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2/models"
)

func fetchEc2RegionalConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	var errs error

	svc := c.Services().Ec2
	var regionalConfig models.RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		c.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = errors.Join(errs, err)
	} else if resp != nil {
		regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId
	}

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
	if err != nil {
		c.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = errors.Join(errs, err)
	}

	if ebsResp != nil && ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}

	res <- regionalConfig

	if errs != nil {
		return fmt.Errorf("error getting regional config: %w", errs)
	}

	return nil
}
