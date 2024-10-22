package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetAlarmsOutput{}
	require.NoError(t, faker.FakeObject(&b))
	b.NextPageToken = nil
	m.EXPECT().GetAlarms(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&b, nil)

	return client.Services{
		Lightsail: m,
	}
}

func TestAlarms(t *testing.T) {
	client.AwsMockTestHelper(t, Alarms(), buildAlarmsMock, client.TestOptions{})
}
