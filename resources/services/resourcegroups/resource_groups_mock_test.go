package resourcegroups

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildResourceGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockResourcegroupsClient(ctrl)
	gId := types.GroupIdentifier{}
	require.NoError(t, faker.FakeObject(&gId))

	groupResponse := types.Group{}
	require.NoError(t, faker.FakeObject(&groupResponse))

	tagsResponse := resourcegroups.GetTagsOutput{}
	require.NoError(t, faker.FakeObject(&tagsResponse))

	query := types.GroupQuery{}
	require.NoError(t, faker.FakeObject(&query))

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.ListGroupsOutput{
			GroupIdentifiers: []types.GroupIdentifier{gId},
		}, nil)
	m.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.GetGroupOutput{
			Group: &groupResponse,
		}, nil)
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagsResponse, nil)
	m.EXPECT().GetGroupQuery(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&resourcegroups.GetGroupQueryOutput{
			GroupQuery: &query,
		}, nil)

	return client.Services{
		Resourcegroups: m,
	}
}

func TestResourceGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceGroups(), buildResourceGroupsMock, client.TestOptions{})
}
