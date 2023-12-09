package lightsail

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
)

func buildAlarmsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetAlarmsOutput{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
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
