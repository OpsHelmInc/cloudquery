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

func buildDisks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var disks lightsail.GetDisksOutput
	require.NoError(t, faker.FakeObject(&disks))

	disks.NextPageToken = nil
	mock.EXPECT().GetDisks(
		gomock.Any(),
		&lightsail.GetDisksInput{},
		gomock.Any(),
	).Return(
		&disks,
		nil,
	)

	var diskSnapshots lightsail.GetDiskSnapshotsOutput
	require.NoError(t, faker.FakeObject(&diskSnapshots))

	diskSnapshots.NextPageToken = nil
	mock.EXPECT().GetDiskSnapshots(
		gomock.Any(),
		&lightsail.GetDiskSnapshotsInput{},
		gomock.Any(),
	).Return(
		&diskSnapshots,
		nil,
	)

	return client.Services{Lightsail: mock}
}

func TestLightsailDisks(t *testing.T) {
	client.AwsMockTestHelper(t, Disks(), buildDisks, client.TestOptions{})
}
