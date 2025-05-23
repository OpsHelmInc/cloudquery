package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/iam/models"
)

func SamlIdentityProviders() *schema.Table {
	tableName := "aws_iam_saml_identity_providers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html`,
		Resolver:            fetchIamSamlIdentityProviders,
		PreResourceResolver: getSamlIdentityProvider,
		Transform: transformers.TransformWithStruct(
			&models.IAMSAMLIdentityProviderWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIamSamlIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.SAMLProviderList
	return nil
}

func getSamlIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	p := resource.Item.(types.SAMLProviderListEntry)

	providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = models.IAMSAMLIdentityProviderWrapper{
		GetSAMLProviderOutput: providerResponse,
		Arn:                   *p.Arn,
		Tags:                  client.TagsToMap(providerResponse.Tags),
	}
	return nil
}
