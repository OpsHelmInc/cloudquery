package kms

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func keyGrants() *schema.Table {
	tableName := "aws_kms_key_grants"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/kms/latest/APIReference/API_GrantListEntry.html`,
		Resolver:    fetchKmsKeyGrants,
		Transform:   transformers.TransformWithStruct(&types.GrantListEntry{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "key_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "grant_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("GrantId"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchKmsKeyGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	k := parent.Item.(*types.KeyMetadata)
	config := kms.ListGrantsInput{
		KeyId: k.Arn,
		Limit: aws.Int32(100),
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceKms).Kms
	p := kms.NewListGrantsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *kms.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Grants
	}
	return nil
}
