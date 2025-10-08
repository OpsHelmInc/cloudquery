package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildIotStreamsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	streams := iot.ListStreamsOutput{}
	err := faker.FakeObject(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.NextToken = nil
	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streams, nil)

	streamOutput := iot.DescribeStreamOutput{}
	err = faker.FakeObject(&streamOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeStream(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&streamOutput, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotStreams(t *testing.T) {
	client.AwsMockTestHelper(t, Streams(), buildIotStreamsMock, client.TestOptions{})
}
