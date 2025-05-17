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

func buildElasticacheReplicationGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	output := elasticache.DescribeReplicationGroupsOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeReplicationGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&output, nil)
	mockElasticache.EXPECT().ListTagsForResource(
		gomock.Any(),
		&elasticache.ListTagsForResourceInput{ResourceName: output.ReplicationGroups[0].ARN},
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

func TestElasticacheReplicationGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ReplicationGroups(), buildElasticacheReplicationGroups, client.TestOptions{})
}
