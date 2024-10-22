package applicationautoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildScheduledActions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationautoscalingClient(ctrl)
	services := client.Services{
		Applicationautoscaling: m,
	}
	c := types.ScheduledAction{}
	require.NoError(t, faker.FakeObject(&c))

	output := &applicationautoscaling.DescribeScheduledActionsOutput{
		ScheduledActions: []types.ScheduledAction{c},
	}
	m.EXPECT().DescribeScheduledActions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		output,
		nil,
	)

	return services
}

func TestScheduledActions(t *testing.T) {
	client.AllNamespaces = []string{"test-namespace"} // Just one

	client.AwsMockTestHelper(t, ScheduledActions(), buildScheduledActions, client.TestOptions{})
}
