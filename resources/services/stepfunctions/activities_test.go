package stepfunctions

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildActivities(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSfnClient(ctrl)
	ali := types.ActivityListItem{}
	require.NoError(t, faker.FakeObject(&ali))

	m.EXPECT().ListActivities(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListActivitiesOutput{
			Activities: []types.ActivityListItem{ali},
		}, nil)

	return client.Services{
		Sfn: m,
	}
}

func TestActivities(t *testing.T) {
	client.AwsMockTestHelper(t, Activities(), buildActivities, client.TestOptions{})
}
