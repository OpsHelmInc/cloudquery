package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/iam/models"
)

func OpenidConnectIdentityProviders() *schema.Table {
	tableName := "aws_iam_openid_connect_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOpenIDConnectProvider.html`,
		Resolver:            fetchIamOpenidConnectIdentityProviders,
		PreResourceResolver: getOpenIdConnectIdentityProvider,
		Transform:           transformers.TransformWithStruct(&models.IamOpenIdIdentityProviderWrapper{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIamOpenidConnectIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	response, err := svc.ListOpenIDConnectProviders(ctx, &iam.ListOpenIDConnectProvidersInput{}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.OpenIDConnectProviderList
	return nil
}

func getOpenIdConnectIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam

	p := resource.Item.(types.OpenIDConnectProviderListEntry)
	providerResponse, err := svc.GetOpenIDConnectProvider(ctx, &iam.GetOpenIDConnectProviderInput{OpenIDConnectProviderArn: p.Arn}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = &models.IamOpenIdIdentityProviderWrapper{GetOpenIDConnectProviderOutput: providerResponse, Arn: *p.Arn}
	return nil
}
