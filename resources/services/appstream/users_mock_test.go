package appstream

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/golang/mock/gomock"
)

func buildAppstreamUsersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.User{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

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
