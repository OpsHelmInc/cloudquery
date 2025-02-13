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

func buildFirewallDomainListMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	fdlm := types.FirewallDomainListMetadata{}
	require.NoError(t, faker.FakeObject(&fdlm))

	m.EXPECT().ListFirewallDomainLists(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListFirewallDomainListsOutput{
			FirewallDomainLists: []types.FirewallDomainListMetadata{fdlm},
		}, nil)

	fdl := types.FirewallDomainList{}
	require.NoError(t, faker.FakeObject(&fdl))

	m.EXPECT().GetFirewallDomainList(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.GetFirewallDomainListOutput{
			FirewallDomainList: &fdl,
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestFirewallDomainList(t *testing.T) {
	client.AwsMockTestHelper(t, FirewallDomainLists(), buildFirewallDomainListMock, client.TestOptions{})
}
