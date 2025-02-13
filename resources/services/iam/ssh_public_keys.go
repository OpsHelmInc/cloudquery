package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/ohaws"
)

func sshPublicKeys() *schema.Table {
	tableName := "aws_iam_ssh_public_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SSHPublicKeyMetadata.html`,
		Resolver:    fetchIamSshPublicKeys,
		Transform:   transformers.TransformWithStruct(&types.SSHPublicKeyMetadata{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "user_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
			{
				Name:                "ssh_public_key_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("SSHPublicKeyId"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchIamSshPublicKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListSSHPublicKeysPaginator(svc, &iam.ListSSHPublicKeysInput{
		UserName: parent.Item.(*ohaws.User).UserName,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SSHPublicKeys
	}
	return nil
}
