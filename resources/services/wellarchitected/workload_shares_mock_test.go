package wellarchitected

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWorkloadShares(t *testing.T, m *mocks.MockWellarchitectedClient, workload *types.Workload) {
	var summary types.WorkloadShareSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListWorkloadShares(gomock.Any(),
		&wellarchitected.ListWorkloadSharesInput{
			WorkloadId: workload.WorkloadId,
			MaxResults: aws.Int32(50),
		},
		gomock.Any()).
		Return(
			&wellarchitected.ListWorkloadSharesOutput{
				WorkloadShareSummaries: []types.WorkloadShareSummary{summary},
			},
			nil,
		)
}
