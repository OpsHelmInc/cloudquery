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

func buildSubscriptions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	subscription := shield.DescribeSubscriptionOutput{}
	require.NoError(t, faker.FakeObject(&subscription))
	m.EXPECT().DescribeSubscription(gomock.Any(), gomock.Any(), gomock.Any()).Return(&subscription, nil)

	return client.Services{
		Shield: m,
	}
}

func TestSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, Subscriptions(), buildSubscriptions, client.TestOptions{})
}
