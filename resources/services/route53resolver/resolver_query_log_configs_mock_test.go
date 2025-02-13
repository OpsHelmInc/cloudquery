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

func buildResolverQueryLogConfigsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	rqlc := types.ResolverQueryLogConfig{}
	require.NoError(t, faker.FakeObject(&rqlc))

	m.EXPECT().ListResolverQueryLogConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListResolverQueryLogConfigsOutput{
			ResolverQueryLogConfigs: []types.ResolverQueryLogConfig{rqlc},
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestResolverQueryLogConfigs(t *testing.T) {
	client.AwsMockTestHelper(t, ResolverQueryLogConfigs(), buildResolverQueryLogConfigsMock, client.TestOptions{})
}
