package scheduler

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSchedulerScheduleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSchedulerClient(ctrl)
	object := types.ScheduleGroupSummary{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListScheduleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&scheduler.ListScheduleGroupsOutput{
			ScheduleGroups: []types.ScheduleGroupSummary{object},
		}, nil)

	tagsOutput := scheduler.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Scheduler: m,
	}
}

func TestSchedulerSchedulerGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ScheduleGroups(), buildSchedulerScheduleGroupsMock, client.TestOptions{})
}
