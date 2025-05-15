package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchAcmCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	var input acm.ListCertificatesInput
	paginator := acm.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.CertificateSummaryList
	}
	return nil
}

func getCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	input := acm.DescribeCertificateInput{CertificateArn: resource.Item.(types.CertificateSummary).CertificateArn}
	output, err := svc.DescribeCertificate(ctx, &input)
	if err != nil {
		return err
	}

	cert := ohaws.ACMCertificate{
		CertificateDetail: *output.Certificate,
	}

	tags, err := getCertificateTags(ctx, svc, *cert.CertificateArn)
	if err != nil {
		return err
	}
	cert.Tags = tags

	resource.Item = &cert
	return nil
}

func getCertificateTags(ctx context.Context, svc services.AcmClient, certificateArn string) ([]types.Tag, error) {
	out, err := svc.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{CertificateArn: &certificateArn})
	if err != nil {
		return nil, err
	}

	return out.Tags, nil
}
