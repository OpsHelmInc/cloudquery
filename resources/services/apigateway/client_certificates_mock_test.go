package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	c := types.ClientCertificate{}
	err := faker.FakeObject(&c)
	if err != nil {
		t.Fatal(err)
	}
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
