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

func buildFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var f types.FileSystem
	require.NoError(t, faker.FakeObject(&f))

	f.FileSystemType = types.FileSystemTypeLustre
	f.Lifecycle = types.FileSystemLifecycleAvailable
	f.StorageType = types.StorageTypeHdd
	m.EXPECT().DescribeFileSystems(
		gomock.Any(),
		&fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeFileSystemsOutput{FileSystems: []types.FileSystem{f}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestFilesystems(t *testing.T) {
	client.AwsMockTestHelper(t, FileSystems(), buildFilesystemsMock, client.TestOptions{})
}
