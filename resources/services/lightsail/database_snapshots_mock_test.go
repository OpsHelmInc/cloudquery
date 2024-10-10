package lightsail

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDatabaseSnapshotsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	s := lightsail.GetRelationalDatabaseSnapshotsOutput{}
	require.NoError(t, faker.FakeObject(&s))
	s.NextPageToken = nil
	m.EXPECT().GetRelationalDatabaseSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestDatabaseSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, DatabaseSnapshots(), buildDatabaseSnapshotsMock, client.TestOptions{})
}
