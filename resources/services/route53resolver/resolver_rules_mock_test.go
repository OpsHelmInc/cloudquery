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

func buildResolverRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	rr := types.ResolverRule{}
	require.NoError(t, faker.FakeObject(&rr))

	m.EXPECT().ListResolverRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListResolverRulesOutput{
			ResolverRules: []types.ResolverRule{rr},
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestResolverRules(t *testing.T) {
	client.AwsMockTestHelper(t, ResolverRules(), buildResolverRulesMock, client.TestOptions{})
}
