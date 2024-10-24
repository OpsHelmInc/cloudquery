package kafka

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildKafkaNodesMock(t *testing.T, m *mocks.MockKafkaClient) {
	object := types.NodeInfo{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().ListNodes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&kafka.ListNodesOutput{
			NodeInfoList: []types.NodeInfo{object},
		}, nil)
}
