package secretsmanager

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func secretVersions() *schema.Table {
	tableName := "aws_secretsmanager_secret_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_ListSecretVersionIds.html`,
		Resolver:    fetchSecretsmanagerSecretsVersions,
		Transform:   transformers.TransformWithStruct(&types.SecretVersionsListEntry{}, transformers.WithPrimaryKeyComponents("VersionId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "secret_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSecretsmanagerSecretsVersions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	secret := resource.Item.(*secretsmanager.DescribeSecretOutput)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSecretsmanager).Secretsmanager
	paginator := secretsmanager.NewListSecretVersionIdsPaginator(svc, &secretsmanager.ListSecretVersionIdsInput{
		SecretId:          secret.ARN,
		IncludeDeprecated: aws.Bool(true),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *secretsmanager.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Versions
	}
	return nil
}
