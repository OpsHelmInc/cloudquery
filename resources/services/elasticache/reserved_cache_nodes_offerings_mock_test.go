package elasticache

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildElasticacheReservedCacheNodesOfferings(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeReservedCacheNodesOfferingsOutput{}
	require.NoError(t, faker.FakeObject(&output))
	output.Marker = nil

	mockElasticache.EXPECT().DescribeReservedCacheNodesOfferings(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheReservedCacheNodesOfferings(t *testing.T) {
	client.AwsMockTestHelper(t, ReservedCacheNodesOfferings(), buildElasticacheReservedCacheNodesOfferings, client.TestOptions{})
}
