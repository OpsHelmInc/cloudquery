package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildModelVersions(t *testing.T, client *mocks.MockFrauddetectorClient) {
	data := types.ModelVersionDetail{}
	require.NoError(t, faker.FakeObject(&data))

	client.EXPECT().DescribeModelVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.DescribeModelVersionsOutput{ModelVersionDetails: []types.ModelVersionDetail{data}}, nil,
	)
}
