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

func buildEventTypes(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.EventType{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetEventTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetEventTypesOutput{EventTypes: []types.EventType{data}}, nil,
	)

	addTagsCall(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestEventTypes(t *testing.T) {
	client.AwsMockTestHelper(t, EventTypes(), buildEventTypes, client.TestOptions{})
}
