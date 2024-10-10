package codepipeline

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWebhooks(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCodepipelineClient(ctrl)

	var webhook types.ListWebhookItem
	require.NoError(t, faker.FakeObject(&webhook))

	mock.EXPECT().ListWebhooks(
		gomock.Any(),
		&codepipeline.ListWebhooksInput{},
		gomock.Any(),
	).Return(
		&codepipeline.ListWebhooksOutput{Webhooks: []types.ListWebhookItem{webhook}},
		nil,
	)

	return client.Services{Codepipeline: mock}
}

func TestCodePipelineWebhooks(t *testing.T) {
	client.AwsMockTestHelper(t, Webhooks(), buildWebhooks, client.TestOptions{})
}
