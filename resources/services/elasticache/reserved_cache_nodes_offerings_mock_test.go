package elasticache

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/golang/mock/gomock"
)

func buildElasticacheReservedCacheNodesOfferings(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeReservedCacheNodesOfferingsOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeReservedCacheNodesOfferings(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheReservedCacheNodesOfferings(t *testing.T) {
	client.AwsMockTestHelper(t, ReservedCacheNodesOfferings(), buildElasticacheReservedCacheNodesOfferings, client.TestOptions{})
}
