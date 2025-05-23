package cloudtrail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCloudtrailChannelsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	channel := types.Channel{}
	require.NoError(t, faker.FakeObject(&channel))

	m.EXPECT().ListChannels(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.ListChannelsOutput{
			Channels: []types.Channel{channel},
		},
		nil,
	)

	channelOutput := &cloudtrail.GetChannelOutput{}
	require.NoError(t, faker.FakeObject(&channelOutput))

	m.EXPECT().GetChannel(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		channelOutput,
		nil,
	)

	return services
}

func TestCloudtrailChannels(t *testing.T) {
	client.AwsMockTestHelper(t, Channels(), buildCloudtrailChannelsMock, client.TestOptions{})
}
