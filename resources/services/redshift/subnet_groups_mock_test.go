package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)

	g := types.ClusterSubnetGroup{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().DescribeClusterSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeClusterSubnetGroupsOutput{
			ClusterSubnetGroups: []types.ClusterSubnetGroup{g},
		}, nil)
	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftSubnetGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SubnetGroups(), buildSubnetGroupsMock, client.TestOptions{})
}
