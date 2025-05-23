package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIotStreamsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	streams := iot.ListStreamsOutput{}
	require.NoError(t, faker.FakeObject(&streams))
	streams.NextToken = nil
	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streams, nil)

	streamOutput := iot.DescribeStreamOutput{}
	require.NoError(t, faker.FakeObject(&streamOutput))
	m.EXPECT().DescribeStream(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streamOutput, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotStreams(t *testing.T) {
	client.AwsMockTestHelper(t, Streams(), buildIotStreamsMock, client.TestOptions{})
}
