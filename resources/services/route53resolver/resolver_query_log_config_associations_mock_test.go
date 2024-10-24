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

func buildResolverQueryLogConfigAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	rqlca := types.ResolverQueryLogConfigAssociation{}
	require.NoError(t, faker.FakeObject(&rqlca))

	m.EXPECT().ListResolverQueryLogConfigAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListResolverQueryLogConfigAssociationsOutput{
			ResolverQueryLogConfigAssociations: []types.ResolverQueryLogConfigAssociation{rqlca},
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestResolverQueryLogConfigAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, ResolverQueryLogConfigAssociations(), buildResolverQueryLogConfigAssociationsMock, client.TestOptions{})
}
