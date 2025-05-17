package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildElasticacheClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeCacheClustersOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheClusters(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)
	mockElasticache.EXPECT().ListTagsForResource(
		gomock.Any(),
		&elasticache.ListTagsForResourceInput{ResourceName: output.CacheClusters[0].ARN},
	).Return(
		&elasticache.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key1"), Value: aws.String("val1")}},
		},
		nil,
	)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildElasticacheClusters, client.TestOptions{})
}
