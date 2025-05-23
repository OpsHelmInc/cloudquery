package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildOrganizationsRoots(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := types.Root{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListRoots(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListRootsOutput{
			Roots: []types.Root{g},
		}, nil)

	tt := make([]types.Tag, 3)
	require.NoError(t, faker.FakeObject(&tt))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListTagsForResourceOutput{
			Tags: tt,
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationsRoots(t *testing.T) {
	client.AwsMockTestHelper(t, Roots(), buildOrganizationsRoots, client.TestOptions{})
}
