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

func buildApigatewayDomainNames(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	dm := types.DomainName{}
	require.NoError(t, faker.FakeObject(&dm))

	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetDomainNamesOutput{
			Items: []types.DomainName{dm},
		}, nil)

	bpm := types.BasePathMapping{}
	require.NoError(t, faker.FakeObject(&bpm))

	m.EXPECT().GetBasePathMappings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetBasePathMappingsOutput{
			Items: []types.BasePathMapping{bpm},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestDomainNames(t *testing.T) {
	client.AwsMockTestHelper(t, DomainNames(), buildApigatewayDomainNames, client.TestOptions{})
}
