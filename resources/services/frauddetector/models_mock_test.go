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

func buildModels(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Model{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	fdClient.EXPECT().GetModels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetModelsOutput{Models: []types.Model{data}}, nil,
	)

	buildModelVersions(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestModels(t *testing.T) {
	client.AwsMockTestHelper(t, Models(), buildModels, client.TestOptions{})
}
