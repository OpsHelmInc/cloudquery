package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildNeptuneDBParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var g types.DBParameterGroup
	require.NoError(t, faker.FakeObject(&g))

	mock.EXPECT().DescribeDBParameterGroups(
		gomock.Any(),
		&neptune.DescribeDBParameterGroupsInput{
			Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
		},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBParameterGroupsOutput{DBParameterGroups: []types.DBParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	var p types.Parameter
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().DescribeDBParameters(
		gomock.Any(),
		&neptune.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneDBParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, DbParameterGroups(), buildNeptuneDBParameterGroups, client.TestOptions{})
}
