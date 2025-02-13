package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildNeptuneGlobalClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var gc types.GlobalCluster
	require.NoError(t, faker.FakeObject(&gc))

	mock.EXPECT().DescribeGlobalClusters(gomock.Any(), &neptune.DescribeGlobalClustersInput{}, gomock.Any()).Return(
		&neptune.DescribeGlobalClustersOutput{GlobalClusters: []types.GlobalCluster{gc}},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneGlobalCluster(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalClusters(), buildNeptuneGlobalClusters, client.TestOptions{})
}
