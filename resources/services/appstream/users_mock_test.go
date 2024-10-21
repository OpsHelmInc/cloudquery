package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAppstreamUsersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.User{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeUsersOutput{
			Users: []types.User{object},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildAppstreamUsersMock, client.TestOptions{})
}
