package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCertificatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var parameterGroups docdb.DescribeCertificatesOutput
	require.NoError(t, faker.FakeObject(&parameterGroups))

	parameterGroups.Marker = nil
	m.EXPECT().DescribeCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&parameterGroups,
		nil,
	)
	return services
}

func TestCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, Certificates(), buildCertificatesMock, client.TestOptions{})
}
