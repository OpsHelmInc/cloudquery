package workspaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDirectories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var directory types.WorkspaceDirectory
	require.NoError(t, faker.FakeObject(&directory))

	mock.EXPECT().DescribeWorkspaceDirectories(
		gomock.Any(),
		&workspaces.DescribeWorkspaceDirectoriesInput{},
		gomock.Any(),
	).Return(
		&workspaces.DescribeWorkspaceDirectoriesOutput{Directories: []types.WorkspaceDirectory{directory}},
		nil,
	)

	return client.Services{Workspaces: mock}
}

func TestWorkspacesDirectories(t *testing.T) {
	client.AwsMockTestHelper(t, Directories(), buildDirectories, client.TestOptions{})
}
