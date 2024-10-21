package signer

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Profiles() *schema.Table {
	tableName := "aws_signer_signing_profiles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/signer/latest/api/API_GetSigningProfile.html`,
		Resolver:            fetchProfiles,
		PreResourceResolver: getProfile,
		Transform:           transformers.TransformWithStruct(&signer.GetSigningProfileOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "signer"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ProfileVersionArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSigner).Signer
	config := signer.ListSigningProfilesInput{}

	paginator := signer.NewListSigningProfilesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *signer.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Profiles
	}
	return nil
}

func getProfile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSigner).Signer
	a := resource.Item.(types.SigningProfile)

	profile, err := svc.GetSigningProfile(ctx, &signer.GetSigningProfileInput{ProfileName: a.ProfileName}, func(o *signer.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = profile
	return nil
}
