package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCloudwatchLogsLogGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudwatchlogsClient(ctrl)
	l := types.LogGroup{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().DescribeLogGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeLogGroupsOutput{
			LogGroups: []types.LogGroup{l},
		}, nil)

	sf := types.SubscriptionFilter{}
	require.NoError(t, faker.FakeObject(&sf))
	m.EXPECT().DescribeSubscriptionFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudwatchlogs.DescribeSubscriptionFiltersOutput{
			SubscriptionFilters: []types.SubscriptionFilter{sf},
		}, nil)

	tags := &cloudwatchlogs.ListTagsLogGroupOutput{}
	require.NoError(t, faker.FakeObject(tags))

	m.EXPECT().ListTagsLogGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(tags, nil)

	dataProtectionPolicy := &cloudwatchlogs.GetDataProtectionPolicyOutput{}
	require.NoError(t, faker.FakeObject(dataProtectionPolicy))

	m.EXPECT().GetDataProtectionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(dataProtectionPolicy, nil)

	return client.Services{
		Cloudwatchlogs: m,
	}
}

func TestCloudwatchlogsLogGroups(t *testing.T) {
	client.AwsMockTestHelper(t, LogGroups(), buildCloudwatchLogsLogGroupsMock, client.TestOptions{})
}
