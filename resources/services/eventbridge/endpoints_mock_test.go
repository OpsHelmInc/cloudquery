package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEventbridgeEndpointsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)
	object := types.Endpoint{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListEndpointsOutput{
			Endpoints: []types.Endpoint{object},
		}, nil)

	tagsOutput := eventbridge.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()
	return client.Services{
		Eventbridge: m,
	}
}

func TestEventbridgeEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, Endpoints(), buildEventbridgeEndpointsMock, client.TestOptions{})
}
