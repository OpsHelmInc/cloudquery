package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildOptionGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var s types.OptionGroup
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeOptionGroups(gomock.Any(), &rds.DescribeOptionGroupsInput{}, gomock.Any()).Return(
		&rds.DescribeOptionGroupsOutput{OptionGroupsList: []types.OptionGroup{s}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: s.OptionGroupArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSOptionGroups(t *testing.T) {
	client.AwsMockTestHelper(t, OptionGroups(), buildOptionGroups, client.TestOptions{})
}
