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

func buildApplicationAutoscalingPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationautoscalingClient(ctrl)
	services := client.Services{
		Applicationautoscaling: m,
	}
	c := types.ScalingPolicy{}
	require.NoError(t, faker.FakeObject(&c))

	output := &applicationautoscaling.DescribeScalingPoliciesOutput{
		ScalingPolicies: []types.ScalingPolicy{c},
	}
	m.EXPECT().DescribeScalingPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		output,
		nil,
	)

	return services
}

func TestApplicationAutoscalingPolicies(t *testing.T) {
	client.AllNamespaces = []string{"test-namespace"} // Just one

	client.AwsMockTestHelper(t, Policies(), buildApplicationAutoscalingPoliciesMock, client.TestOptions{})
}
