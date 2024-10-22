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

func buildStaticIps(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockLightsailClient(ctrl)

	var ips lightsail.GetStaticIpsOutput
	require.NoError(t, faker.FakeObject(&ips))

	ips.NextPageToken = nil

	mock.EXPECT().GetStaticIps(gomock.Any(), &lightsail.GetStaticIpsInput{}, gomock.Any()).Return(&ips, nil)

	return client.Services{Lightsail: mock}
}

func TestStaticIps(t *testing.T) {
	client.AwsMockTestHelper(t, StaticIps(), buildStaticIps, client.TestOptions{})
}
