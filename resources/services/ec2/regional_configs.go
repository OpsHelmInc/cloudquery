package ec2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2/models"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RegionalConfigs() *schema.Table {
	tableName := "aws_ec2_regional_configs"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsDefaultKmsKeyId.html
https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetEbsEncryptionByDefault.html`,
		Resolver:  fetchEc2RegionalConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform: transformers.TransformWithStruct(&models.RegionalConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveEc2RegionalConfigArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchEc2RegionalConfigs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	svc := cl.Services(client.AWSServiceEc2).Ec2
	var regionalConfig models.RegionalConfig
	resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId

	ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{}, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	if ebsResp.EbsEncryptionByDefault != nil {
		regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
	}
	res <- regionalConfig
	return nil
}

func resolveEc2RegionalConfigArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "regional-config",
	}
	return resource.Set(c.Name, a.String())
}
