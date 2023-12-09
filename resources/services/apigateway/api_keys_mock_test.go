package apigateway

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/golang/mock/gomock"
)

func buildApiKeysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	a := types.ApiKey{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetApiKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetApiKeysOutput{
			Items: []types.ApiKey{a},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestAPIKeys(t *testing.T) {
	client.AwsMockTestHelper(t, ApiKeys(), buildApiKeysMock, client.TestOptions{})
}
