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

func buildRDSEventSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var s types.EventSubscription
	require.NoError(t, faker.FakeObject(&s))

	mock.EXPECT().DescribeEventSubscriptions(gomock.Any(), &rds.DescribeEventSubscriptionsInput{}, gomock.Any()).Return(
		&rds.DescribeEventSubscriptionsOutput{EventSubscriptionsList: []types.EventSubscription{s}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: s.EventSubscriptionArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, EventSubscriptions(), buildRDSEventSubscriptions, client.TestOptions{})
}
