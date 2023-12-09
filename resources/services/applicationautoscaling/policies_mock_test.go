package applicationautoscaling

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/golang/mock/gomock"
)

func buildApplicationAutoscalingPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApplicationautoscalingClient(ctrl)
	services := client.Services{
		Applicationautoscaling: m,
	}
	c := types.ScalingPolicy{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
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
