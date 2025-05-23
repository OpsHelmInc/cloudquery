package appconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDeploymentStrategies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppconfigClient(ctrl)

	var ds types.DeploymentStrategy
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().ListDeploymentStrategies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appconfig.ListDeploymentStrategiesOutput{
			Items: []types.DeploymentStrategy{ds},
		},
		nil,
	)

	return client.Services{Appconfig: m}
}

func TestDeploymentStrategies(t *testing.T) {
	client.AwsMockTestHelper(t, DeploymentStrategies(), buildDeploymentStrategies, client.TestOptions{})
}
