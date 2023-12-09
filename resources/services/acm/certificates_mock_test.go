package acm

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/golang/mock/gomock"
)

func buildACMCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAcmClient(ctrl)

	var cs types.CertificateSummary
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListCertificates(
		gomock.Any(),
		&acm.ListCertificatesInput{},
		gomock.Any(),
	).Return(
		&acm.ListCertificatesOutput{CertificateSummaryList: []types.CertificateSummary{cs}},
		nil,
	)

	var cert types.CertificateDetail
	if err := faker.FakeObject(&cert); err != nil {
		t.Fatal(err)
	}
	cert.CertificateArn = cs.CertificateArn
	mock.EXPECT().DescribeCertificate(
		gomock.Any(),
		&acm.DescribeCertificateInput{CertificateArn: cs.CertificateArn},
		gomock.Any(),
	).Return(
		&acm.DescribeCertificateOutput{Certificate: &cert},
		nil,
	)

	mock.EXPECT().ListTagsForCertificate(
		gomock.Any(),
		&acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn},
	).Return(
		&acm.ListTagsForCertificateOutput{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
	return client.Services{Acm: mock}
}

func TestACMCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildACMCertificates, client.TestOptions{})
}
