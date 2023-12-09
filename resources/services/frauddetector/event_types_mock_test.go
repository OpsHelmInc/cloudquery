package frauddetector

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/golang/mock/gomock"
)

func buildEventTypes(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.EventType{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

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
