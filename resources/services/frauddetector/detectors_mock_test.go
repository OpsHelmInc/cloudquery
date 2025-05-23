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

func buildDetectors(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Detector{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetDetectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetDetectorsOutput{Detectors: []types.Detector{data}}, nil,
	)

	buildRules(t, fdClient)
	addTagsCall(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestDetectors(t *testing.T) {
	client.AwsMockTestHelper(t, Detectors(), buildDetectors, client.TestOptions{})
}
