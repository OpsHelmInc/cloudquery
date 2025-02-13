package wellarchitected

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildWorkloadsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWellarchitectedClient(ctrl)

	var summary types.WorkloadSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListWorkloads(gomock.Any(),
		&wellarchitected.ListWorkloadsInput{MaxResults: aws.Int32(50)},
		gomock.Any()).
		Return(&wellarchitected.ListWorkloadsOutput{WorkloadSummaries: []types.WorkloadSummary{summary}}, nil)

	var workload types.Workload
	require.NoError(t, faker.FakeObject(&workload))

	m.EXPECT().GetWorkload(gomock.Any(), &wellarchitected.GetWorkloadInput{WorkloadId: summary.WorkloadId}, gomock.Any()).
		Return(&wellarchitected.GetWorkloadOutput{Workload: &workload}, nil)

	buildWorkloadMilestones(t, m, &workload)
	buildWorkloadShares(t, m, &workload)

	return client.Services{Wellarchitected: m}
}

func TestWorkloads(t *testing.T) {
	client.AwsMockTestHelper(t, Workloads(), buildWorkloadsMock, client.TestOptions{})
}
