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

func buildStorageVmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var vm types.StorageVirtualMachine
	require.NoError(t, faker.FakeObject(&vm))
	m.EXPECT().DescribeStorageVirtualMachines(
		gomock.Any(),
		&fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeStorageVirtualMachinesOutput{StorageVirtualMachines: []types.StorageVirtualMachine{vm}},
		nil,
	)
	return client.Services{
		Fsx: m,
	}
}

func TestStorageVms(t *testing.T) {
	client.AwsMockTestHelper(t, StorageVirtualMachines(), buildStorageVmsMock, client.TestOptions{})
}
