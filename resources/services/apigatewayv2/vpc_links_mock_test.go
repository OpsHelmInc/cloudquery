package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApigatewayv2VpcLinks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	v := types.VpcLink{}
	require.NoError(t, faker.FakeObject(&v))
	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apigatewayv2.GetVpcLinksOutput{
			Items: []types.VpcLink{v},
		}, nil)

	return client.Services{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2VpcLinks(t *testing.T) {
	client.AwsMockTestHelper(t, VpcLinks(), buildApigatewayv2VpcLinks, client.TestOptions{})
}
