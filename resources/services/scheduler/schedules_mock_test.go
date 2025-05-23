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

func buildSchedulerSchedulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSchedulerClient(ctrl)
	object := types.ScheduleSummary{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListSchedules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&scheduler.ListSchedulesOutput{
			Schedules: []types.ScheduleSummary{object},
		}, nil)
	object2 := scheduler.GetScheduleOutput{}
	require.NoError(t, faker.FakeObject(&object2))

	m.EXPECT().GetSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&object2, nil)

	tagsOutput := scheduler.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Scheduler: m,
	}
}

func TestSchedulesMock(t *testing.T) {
	client.AwsMockTestHelper(t, Schedules(), buildSchedulerSchedulesMock, client.TestOptions{})
}
