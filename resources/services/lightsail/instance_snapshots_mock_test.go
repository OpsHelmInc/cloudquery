package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildInstanceSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var is lightsail.GetInstanceSnapshotsOutput
	require.NoError(t, faker.FakeObject(&is))

	is.NextPageToken = nil

	mock.EXPECT().GetInstanceSnapshots(
		gomock.Any(),
		&lightsail.GetInstanceSnapshotsInput{},
		gomock.Any(),
	).Return(&is, nil)

	return client.Services{Lightsail: mock}
}

func TestLightsailInstanceSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceSnapshots(), buildInstanceSnapshots, client.TestOptions{})
}
