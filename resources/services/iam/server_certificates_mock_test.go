package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIamServerCerts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.ServerCertificateMetadata{}
	require.NoError(t, faker.FakeObject(&u))

	m.EXPECT().ListServerCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListServerCertificatesOutput{
			ServerCertificateMetadataList: []iamTypes.ServerCertificateMetadata{u},
		}, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamServerCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, ServerCertificates(), buildIamServerCerts, client.TestOptions{})
}
