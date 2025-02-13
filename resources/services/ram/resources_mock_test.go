package ram

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRamResourcesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.Resource{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.ListResourcesOutput{
			Resources: []types.Resource{object},
		}, nil).MinTimes(1)

	return client.Services{
		Ram: m,
	}
}

func TestRamResources(t *testing.T) {
	client.AwsMockTestHelper(t, Resources(), buildRamResourcesMock, client.TestOptions{})
}
