package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApprunnerAutoScalingConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApprunnerClient(ctrl)
	as := types.AutoScalingConfiguration{}
	require.NoError(t, faker.FakeObject(&as))

	m.EXPECT().ListAutoScalingConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListAutoScalingConfigurationsOutput{
			AutoScalingConfigurationSummaryList: []types.AutoScalingConfigurationSummary{
				{AutoScalingConfigurationArn: as.AutoScalingConfigurationArn},
			},
		}, nil)

	m.EXPECT().DescribeAutoScalingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.DescribeAutoScalingConfigurationOutput{
			AutoScalingConfiguration: &as,
		}, nil)
	tags := types.Tag{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)
	return client.Services{
		Apprunner: m,
	}
}

func TestApprunnerAutoScalingConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, AutoScalingConfigurations(), buildApprunnerAutoScalingConfigurationsMock, client.TestOptions{})
}
