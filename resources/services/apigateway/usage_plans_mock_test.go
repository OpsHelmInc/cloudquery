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

func buildApigatewayUsagePlans(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	u := types.UsagePlan{}
	require.NoError(t, faker.FakeObject(&u))

	m.EXPECT().GetUsagePlans(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetUsagePlansOutput{
			Items: []types.UsagePlan{u},
		}, nil)

	uk := types.UsagePlanKey{}
	require.NoError(t, faker.FakeObject(&uk))

	m.EXPECT().GetUsagePlanKeys(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetUsagePlanKeysOutput{
			Items: []types.UsagePlanKey{uk},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestUsagePlans(t *testing.T) {
	client.AwsMockTestHelper(t, UsagePlans(), buildApigatewayUsagePlans, client.TestOptions{})
}
