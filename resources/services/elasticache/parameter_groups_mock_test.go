package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildElasticacheParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mockElasticache := mocks.NewMockElasticacheClient(ctrl)
	parameterGroupsOutput := elasticache.DescribeCacheParameterGroupsOutput{}
	require.NoError(t, faker.FakeObject(&parameterGroupsOutput))
	parameterGroupsOutput.Marker = nil

	parametersOutput := elasticache.DescribeCacheParametersOutput{}
	require.NoError(t, faker.FakeObject(&parametersOutput))
	parametersOutput.Marker = nil

	mockElasticache.EXPECT().DescribeCacheParameterGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(&parameterGroupsOutput, nil)

	return client.Services{
		Elasticache: mockElasticache,
	}
}

func TestElasticacheParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ParameterGroups(), buildElasticacheParameterGroups, client.TestOptions{})
}
