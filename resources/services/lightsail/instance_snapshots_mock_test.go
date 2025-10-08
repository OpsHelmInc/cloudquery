package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildInstanceSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var is lightsail.GetInstanceSnapshotsOutput
	if err := faker.FakeObject(&is); err != nil {
		t.Fatal(err)
	}
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
