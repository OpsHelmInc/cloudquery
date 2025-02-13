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

func buildApiKeysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	a := types.ApiKey{}
	require.NoError(t, faker.FakeObject(&a))

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
