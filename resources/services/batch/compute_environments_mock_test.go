package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildBatchComputeEnvironmentsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockBatchClient(ctrl)
	services := client.Services{
		Batch: m,
	}
	a := types.ComputeEnvironmentDetail{}
	require.NoError(t, faker.FakeObject(&a))

	m.EXPECT().DescribeComputeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&batch.DescribeComputeEnvironmentsOutput{
			ComputeEnvironments: []types.ComputeEnvironmentDetail{a},
		}, nil)

	tagResponse := batch.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return services
}

func TestBatchComputeEnvironments(t *testing.T) {
	client.AwsMockTestHelper(t, ComputeEnvironments(), buildBatchComputeEnvironmentsMock, client.TestOptions{})
}
