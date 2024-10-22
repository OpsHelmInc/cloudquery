package amplify

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAmplifyClient(ctrl)

	var app types.App
	require.NoError(t, faker.FakeObject(&app))

	m.EXPECT().ListApps(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&amplify.ListAppsOutput{
			Apps: []types.App{app},
		},
		nil,
	)

	return client.Services{Amplify: m}
}

func TestApps(t *testing.T) {
	client.AwsMockTestHelper(t, Apps(), buildApps, client.TestOptions{})
}
