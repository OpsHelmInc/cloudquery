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

func buildApigatewayVpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayClient(ctrl)

	l := types.VpcLink{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigateway.GetVpcLinksOutput{
			Items: []types.VpcLink{l},
		}, nil)

	return client.Services{
		Apigateway: m,
	}
}

func TestVpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, VpcLinks(), buildApigatewayVpcLinks, client.TestOptions{})
}
