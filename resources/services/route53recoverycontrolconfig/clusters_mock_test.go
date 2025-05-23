package route53recoverycontrolconfig

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoverycontrolconfigClient(ctrl)

	var c types.Cluster
	require.NoError(t, faker.FakeObject(&c))

	m.EXPECT().ListClusters(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&route53recoverycontrolconfig.ListClustersOutput{
			Clusters: []types.Cluster{c},
		},
		nil,
	)

	return client.Services{
		Route53recoverycontrolconfig: m,
	}
}

func TestClusters(t *testing.T) {
	client.AwsMockTestHelper(t, Clusters(), buildClusters, client.TestOptions{Region: "us-west-2"})
}
