package frauddetector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func addTagsCall(t *testing.T, client *mocks.MockFrauddetectorClient) {
	var data []types.Tag
	require.NoError(t, faker.FakeObject(&data))

	client.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.ListTagsForResourceOutput{Tags: data}, nil,
	)
}
