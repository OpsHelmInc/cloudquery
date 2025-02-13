package docdb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ev docdb.DescribeDBSubnetGroupsOutput
	require.NoError(t, faker.FakeObject(&ev))

	ev.Marker = nil
	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ev,
		nil,
	)

	var tags docdb.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags,
		nil,
	)

	return services
}

func TestSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildSubnetGroupsMock, client.TestOptions{})
}
