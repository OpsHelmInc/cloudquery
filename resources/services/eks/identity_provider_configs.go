package eks

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func identityProviderConfigs() *schema.Table {
	tableName := "aws_eks_cluster_oidc_identity_provider_configs"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_OidcIdentityProviderConfig.html`,
		Resolver:            fetchIdentityProviderConfigs,
		PreResourceResolver: getIdentityProviderConfigs,
		Transform:           transformers.TransformWithStruct(&types.OidcIdentityProviderConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("IdentityProviderConfigArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "cluster_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIdentityProviderConfigs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEks).Eks
	paginator := eks.NewListIdentityProviderConfigsPaginator(svc, &eks.ListIdentityProviderConfigsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.IdentityProviderConfigs
	}
	return nil
}

func getIdentityProviderConfigs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEks).Eks
	ipc := resource.Item.(types.IdentityProviderConfig)
	if aws.ToString(ipc.Type) != "oidc" {
		return nil
	}
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeIdentityProviderConfig(
		ctx, &eks.DescribeIdentityProviderConfigInput{
			ClusterName:            cluster.Name,
			IdentityProviderConfig: &ipc}, func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.IdentityProviderConfig.Oidc
	return nil
}
