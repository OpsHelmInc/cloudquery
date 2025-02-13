package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApprunnerVpcConnectorsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApprunnerClient(ctrl)
	vc := types.VpcConnector{}
	require.NoError(t, faker.FakeObject(&vc))

	m.EXPECT().ListVpcConnectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListVpcConnectorsOutput{
			VpcConnectors: []types.VpcConnector{vc},
		}, nil)
	tags := types.Tag{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)

	return client.Services{
		Apprunner: m,
	}
}

func TestApprunnerVpcConnector(t *testing.T) {
	client.AwsMockTestHelper(t, VpcConnectors(), buildApprunnerVpcConnectorsMock, client.TestOptions{})
}
