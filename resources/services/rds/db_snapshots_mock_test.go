package rds

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
)

func buildRDSClient(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBSnapshot
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshots(
		gomock.Any(),
		&rds.DescribeDBSnapshotsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotsOutput{DBSnapshots: []types.DBSnapshot{s}},
		nil,
	)

	var attrs []types.DBSnapshotAttribute
	if err := faker.FakeObject(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		gomock.Any(),
	).Return(
		&rds.DescribeDBSnapshotAttributesOutput{
			DBSnapshotAttributesResult: &types.DBSnapshotAttributesResult{DBSnapshotAttributes: attrs},
		},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSDBSnapshots(t *testing.T) {
	client.AwsMockTestHelper(t, DbSnapshots(), buildRDSClient, client.TestOptions{})
}
