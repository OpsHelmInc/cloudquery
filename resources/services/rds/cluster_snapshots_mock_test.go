package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRDSClientForClusterSnapshots(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBClusterSnapshot
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeDBClusterSnapshots(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterSnapshotsOutput{DBClusterSnapshots: []types.DBClusterSnapshot{s}},
		nil,
	)

	var attrs []types.DBClusterSnapshotAttribute
	require.NoError(t, faker.FakeObject(&attrs))

	mock.EXPECT().DescribeDBClusterSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterSnapshotAttributesOutput{
			DBClusterSnapshotAttributesResult: &types.DBClusterSnapshotAttributesResult{DBClusterSnapshotAttributes: attrs},
		},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSDBClusterSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterSnapshots(), buildRDSClientForClusterSnapshots, client.TestOptions{})
}
