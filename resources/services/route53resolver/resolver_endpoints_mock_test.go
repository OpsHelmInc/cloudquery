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

func buildEndpointsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	re := types.ResolverEndpoint{}
	require.NoError(t, faker.FakeObject(&re))

	m.EXPECT().ListResolverEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListResolverEndpointsOutput{
			ResolverEndpoints: []types.ResolverEndpoint{re},
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}

func TestEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, ResolverEndpoints(), buildEndpointsMock, client.TestOptions{})
}
