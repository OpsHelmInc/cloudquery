package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	c := types.ClientCertificate{}
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().GetClientCertificates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetClientCertificatesOutput{
			Items: []types.ClientCertificate{c},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestClientCertificates(t *testing.T) {
	client.AwsMockTestHelper(t, ClientCertificates(), buildApigatewayClientCertificates, client.TestOptions{})
}
