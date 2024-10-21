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

func buildVolumesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockFsxClient(ctrl)

	var v types.Volume
	require.NoError(t, faker.FakeObject(&v))

	v.Lifecycle = types.VolumeLifecycleAvailable
	v.VolumeType = types.VolumeTypeOntap
	m.EXPECT().DescribeVolumes(
		gomock.Any(),
		&fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&fsx.DescribeVolumesOutput{Volumes: []types.Volume{v}},
		nil,
	)

	return client.Services{
		Fsx: m,
	}
}

func TestVolumes(t *testing.T) {
	client.AwsMockTestHelper(t, Volumes(), buildVolumesMock, client.TestOptions{})
}
