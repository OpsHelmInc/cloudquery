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

func buildVariables(t *testing.T, ctrl *gomock.Controller) client.Services {
	fdClient := mocks.NewMockFrauddetectorClient(ctrl)

	data := types.Variable{}
	require.NoError(t, faker.FakeObject(&data))

	fdClient.EXPECT().GetVariables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetVariablesOutput{Variables: []types.Variable{data}}, nil,
	)

	addTagsCall(t, fdClient)

	return client.Services{
		Frauddetector: fdClient,
	}
}

func TestVariables(t *testing.T) {
	client.AwsMockTestHelper(t, Variables(), buildVariables, client.TestOptions{})
}
