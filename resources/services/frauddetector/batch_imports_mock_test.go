package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildBatchImports(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.BatchImport{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetBatchImportJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetBatchImportJobsOutput{BatchImports: []types.BatchImport{data}}, nil,
	)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestBatchImports(t *testing.T) {
	client.AwsMockTestHelper(t, BatchImports(), buildBatchImports, client.TestOptions{})
}
