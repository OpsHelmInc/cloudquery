package route53resolver

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildFirewallRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	frgm := types.FirewallRuleGroupMetadata{}
	require.NoError(t, faker.FakeObject(&frgm))

	m.EXPECT().ListFirewallRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListFirewallRuleGroupsOutput{
			FirewallRuleGroups: []types.FirewallRuleGroupMetadata{frgm},
		}, nil)

	frg := types.FirewallRuleGroup{}
	require.NoError(t, faker.FakeObject(&frg))

	m.EXPECT().GetFirewallRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.GetFirewallRuleGroupOutput{
			FirewallRuleGroup: &frg,
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestFirewallRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, FirewallRuleGroups(), buildFirewallRuleGroupsMock, client.TestOptions{})
}
