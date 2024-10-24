package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildProtectionGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	pp := shield.ListProtectionGroupsOutput{}
	require.NoError(t, faker.FakeObject(&pp))
	pp.NextToken = nil
	m.EXPECT().ListProtectionGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&pp, nil)

	tags := shield.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Shield: m,
	}
}

func TestProtectionGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ProtectionGroups(), buildProtectionGroups, client.TestOptions{})
}
