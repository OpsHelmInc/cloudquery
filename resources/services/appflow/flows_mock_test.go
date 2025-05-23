package appflow

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appflow"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildFlows(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAppflowClient(ctrl)

	var fd types.FlowDefinition
	require.NoError(t, faker.FakeObject(&fd))

	mock.EXPECT().ListFlows(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&appflow.ListFlowsOutput{
			Flows: []types.FlowDefinition{fd},
		},
		nil,
	)

	var flowOut appflow.DescribeFlowOutput
	require.NoError(t, faker.FakeObject(&flowOut))

	mock.EXPECT().DescribeFlow(
		gomock.Any(),
		&appflow.DescribeFlowInput{FlowName: flowOut.FlowName},
		gomock.Any(),
	).Return(
		&flowOut,
		nil,
	)

	return client.Services{Appflow: mock}
}

func TestFlows(t *testing.T) {
	client.AwsMockTestHelper(t, Flows(), buildFlows, client.TestOptions{})
}
