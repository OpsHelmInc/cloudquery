package acmpca

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	cqtypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func CertificateAuthorities() *schema.Table {
	tableName := "aws_acmpca_certificate_authorities"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthority.html`,
		Resolver:    fetchAcmpcaCertificateAuthorities,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "acm-pca"),
		Transform:   transformers.TransformWithStruct(&types.CertificateAuthority{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     cqtypes.ExtensionTypes.JSON,
				Resolver: resolveCertificateAuthorityTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchAcmpcaCertificateAuthorities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAcmpca).Acmpca
	paginator := acmpca.NewListCertificateAuthoritiesPaginator(svc, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *acmpca.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.CertificateAuthorities
	}
	return nil
}

func resolveCertificateAuthorityTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	certAuthority := resource.Item.(types.CertificateAuthority)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAcmpca).Acmpca
	out, err := svc.ListTags(ctx,
		&acmpca.ListTagsInput{CertificateAuthorityArn: certAuthority.Arn},
		func(o *acmpca.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.Tags))
}
