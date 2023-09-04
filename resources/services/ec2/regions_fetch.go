package ec2

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2/models"
)

func fetchEc2Regions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	output, err := c.Services().Ec2.DescribeRegions(ctx, &ec2.DescribeRegionsInput{AllRegions: aws.Bool(true)})
	if err != nil {
		return err
	}

	for _, r := range output.Regions {
		res <- &models.Region{Region: r}
	}
	return nil
}

func resolveRegionEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	region := resource.Item.(*models.Region)
	switch *region.OptInStatus {
	case "opt-in-not-required", "opted-in":
		return resource.Set(c.Name, true)
	case "not-opted-in":
		return resource.Set(c.Name, false)
	}
	return nil
}

func resolveRegionPartition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, cl.Partition)
}

func resolveRegionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "region",
	}
	return resource.Set(c.Name, a.String())
}

func resolveEc2RegionConfig(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ec2

	var (
		regionalConfig models.RegionalConfig
		errs           error
	)

	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
	if err != nil {
		cl.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = errors.Join(errs, err)
	} else if resp != nil {
		regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId
	}

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
	if err != nil {
		cl.Logger().Error().Err(err).Msg("error calling GetEbsDefaultKmsKeyId")
		errs = errors.Join(errs, err)
	}

	if ebsResp != nil && ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}

	conf := resource.Item.(*models.Region)
	conf.EC2Config = regionalConfig

	err = resource.Set(c.Name, regionalConfig)
	errs = errors.Join(errs, err)

	return errs
}
