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

func buildRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	rgm := types.RuleGroupMetadata{}
	require.NoError(t, faker.FakeObject(&rgm))

	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupMetadata{rgm},
		}, nil)

	rg := networkfirewall.DescribeRuleGroupOutput{}
	require.NoError(t, faker.FakeObject(&rg))

	m.EXPECT().DescribeRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rg, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildRuleGroupsMock, client.TestOptions{})
}
