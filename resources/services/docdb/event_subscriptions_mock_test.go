package docdb

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDocdbClient(ctrl)
	services := client.Services{
		Docdb: m,
	}
	var ec docdb.DescribeEventSubscriptionsOutput
	require.NoError(t, faker.FakeObject(&ec))

	ec.Marker = nil
	m.EXPECT().DescribeEventSubscriptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec,
		nil,
	)

	return services
}

func TestEventSubscriptions(t *testing.T) {
	client.AwsMockTestHelper(t, EventSubscriptions(), buildEventSubscriptionsMock, client.TestOptions{})
}
