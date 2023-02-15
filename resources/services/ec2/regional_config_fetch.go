package ec2

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2/models"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2RegionalConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	var errs []error

	svc := c.Services().Ec2
	var regionalConfig models.RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		c.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = append(errs, err)
	} else if resp != nil {
		regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId
	}

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
	if err != nil {
		c.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = append(errs, err)
	}

	if ebsResp != nil && ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}

	res <- regionalConfig

	// TODO(kyle): errors.Join when we update to go 1.20
	if len(errs) == 1 {
		return errs[0]
	} else if len(errs) > 1 {
		// right now the length can only be 0, 1, or 2... so, concat the errors
		return fmt.Errorf("errors getting regional config: %s, %s", errs[0], errs[1])
	}

	return nil
}
