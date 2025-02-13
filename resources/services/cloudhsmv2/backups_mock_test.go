package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildHSMBackups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudhsmv2Client(ctrl)

	var backups []types.Backup
	require.NoError(t, faker.FakeObject(&backups))

	mock.EXPECT().DescribeBackups(
		gomock.Any(),
		&cloudhsmv2.DescribeBackupsInput{},
		gomock.Any(),
	).Return(
		&cloudhsmv2.DescribeBackupsOutput{Backups: backups},
		nil,
	)

	return client.Services{Cloudhsmv2: mock}
}

func TestBackups(t *testing.T) {
	client.AwsMockTestHelper(t, Backups(), buildHSMBackups, client.TestOptions{})
}
