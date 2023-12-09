package workspaces

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/golang/mock/gomock"
)

func buildDirectories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var directory types.WorkspaceDirectory
	if err := faker.FakeObject(&directory); err != nil {
		t.Fatal(err)
	}

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
