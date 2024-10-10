package backup

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildProtectedResourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBackupClient(ctrl)

	var pr types.ProtectedResource
	require.NoError(t, faker.FakeObject(&pr))

	m.EXPECT().ListProtectedResources(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&backup.ListProtectedResourcesOutput{
			Results: []types.ProtectedResource{pr},
		},
		nil,
	)

	return client.Services{
		Backup: m,
	}
}

func TestProtectedResources(t *testing.T) {
	client.AwsMockTestHelper(t, ProtectedResources(), buildProtectedResourcesMock, client.TestOptions{})
}
