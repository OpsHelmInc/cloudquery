package glacier

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func vaultLockPolicies() *schema.Table {
	tableName := "aws_glacier_vault_lock_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetVaultLock.html`,
		Resolver:    fetchGlacierVaultLockPolicies,
		Transform:   transformers.TransformWithStruct(&glacier.GetVaultLockOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "vault_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Policy"),
			},
		},
	}
}

func fetchGlacierVaultLockPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlacier).Glacier
	p := parent.Item.(types.DescribeVaultOutput)

	response, err := svc.GetVaultLock(ctx, &glacier.GetVaultLockInput{
		VaultName: p.VaultName,
	}, func(options *glacier.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response
	return nil
}
