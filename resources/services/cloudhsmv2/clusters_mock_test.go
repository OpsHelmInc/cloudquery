package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildHSMClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudhsmv2Client(ctrl)

	var clusters []types.Cluster
	require.NoError(t, faker.FakeObject(&clusters))

	mock.EXPECT().DescribeClusters(
		gomock.Any(),
		&cloudhsmv2.DescribeClustersInput{},
		gomock.Any(),
	).Return(
		&cloudhsmv2.DescribeClustersOutput{Clusters: clusters},
		nil,
	)

	return client.Services{Cloudhsmv2: mock}
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildHSMClusters, client.TestOptions{})
}
