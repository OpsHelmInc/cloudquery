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

func buildExternalModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.ExternalModel{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetExternalModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetExternalModelsOutput{ExternalModels: []types.ExternalModel{data}}, nil,
	)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestExternalModels(t *testing.T) {
	client.AwsMockTestHelper(t, ExternalModels(), buildExternalModels, client.TestOptions{})
}
