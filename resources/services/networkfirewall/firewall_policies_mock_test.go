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

func buildFirewallPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	fpm := types.FirewallPolicyMetadata{}
	require.NoError(t, faker.FakeObject(&fpm))

	m.EXPECT().ListFirewallPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListFirewallPoliciesOutput{
			FirewallPolicies: []types.FirewallPolicyMetadata{fpm},
		}, nil)

	fp := networkfirewall.DescribeFirewallPolicyOutput{}
	require.NoError(t, faker.FakeObject(&fp))

	m.EXPECT().DescribeFirewallPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&fp, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestFirewallPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, FirewallPolicies(), buildFirewallPoliciesMock, client.TestOptions{})
}
