package efs

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/golang/mock/gomock"
)

func buildEfsFilesystemsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.FileSystemDescription{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeFileSystems(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&efs.DescribeFileSystemsOutput{
			FileSystems: []types.FileSystemDescription{l},
		}, nil)

	b := efs.DescribeBackupPolicyOutput{}
	err = faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeBackupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	return client.Services{
		Efs: m,
	}
}

func TestEfsFilesystems(t *testing.T) {
	client.AwsMockTestHelper(t, Filesystems(), buildEfsFilesystemsMock, client.TestOptions{})
}
