package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDataRepoTasksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var task types.DataRepositoryTask
	require.NoError(t, faker.FakeObject(&task))
	m.EXPECT().DescribeDataRepositoryTasks(
		gomock.Any(),
		&fsx.DescribeDataRepositoryTasksInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeDataRepositoryTasksOutput{DataRepositoryTasks: []types.DataRepositoryTask{task}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestDataRepoTasks(t *testing.T) {
	client.AwsMockTestHelper(t, DataRepositoryTasks(), buildDataRepoTasksMock, client.TestOptions{})
}
