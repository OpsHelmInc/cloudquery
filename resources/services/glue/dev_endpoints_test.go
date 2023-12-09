package glue

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDevEndpointsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var devEndpoint glue.GetDevEndpointsOutput
	require.NoError(t, faker.FakeObject(&devEndpoint))
	devEndpoint.NextToken = nil
	m.EXPECT().GetDevEndpoints(
		gomock.Any(),
		&glue.GetDevEndpointsInput{},
	).Return(&devEndpoint, nil)

	m.EXPECT().GetTags(
		gomock.Any(),
		gomock.Any(),
	).Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return client.Services{
		Glue: m,
	}
}

func TestDevEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, DevEndpoints(), buildDevEndpointsMock, client.TestOptions{})
}
