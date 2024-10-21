package networkfirewall

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildFirewallsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	fm := types.FirewallMetadata{}
	require.NoError(t, faker.FakeObject(&fm))

	m.EXPECT().ListFirewalls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListFirewallsOutput{
			Firewalls: []types.FirewallMetadata{fm},
		}, nil)

	fo := networkfirewall.DescribeFirewallOutput{}
	require.NoError(t, faker.FakeObject(&fo))

	m.EXPECT().DescribeFirewall(gomock.Any(), gomock.Any(), gomock.Any()).Return(&fo, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestFirewalls(t *testing.T) {
	client.AwsMockTestHelper(t, Firewalls(), buildFirewallsMock, client.TestOptions{})
}
