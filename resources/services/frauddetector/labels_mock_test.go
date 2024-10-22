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

func buildLabels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Label{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetLabels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetLabelsOutput{Labels: []types.Label{data}}, nil,
	)

	addTagsCall(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestLabels(t *testing.T) {
	client.AwsMockTestHelper(t, Labels(), buildLabels, client.TestOptions{})
}
