package codepipeline

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildPipelines(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCodepipelineClient(ctrl)

	var pipeSummary types.PipelineSummary
	require.NoError(t, faker.FakeObject(&pipeSummary))

	mock.EXPECT().ListPipelines(
		gomock.Any(),
		&codepipeline.ListPipelinesInput{},
		gomock.Any(),
	).Return(
		&codepipeline.ListPipelinesOutput{Pipelines: []types.PipelineSummary{pipeSummary}},
		nil,
	)

	var resource codepipeline.GetPipelineOutput
	require.NoError(t, faker.FakeObject(&resource))

	mock.EXPECT().GetPipeline(
		gomock.Any(),
		&codepipeline.GetPipelineInput{Name: pipeSummary.Name},
		gomock.Any(),
	).Return(
		&resource,
		nil,
	)

	tags := &codepipeline.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	mock.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)

	return client.Services{Codepipeline: mock}
}

func TestCodePipelinePipelines(t *testing.T) {
	client.AwsMockTestHelper(t, Pipelines(), buildPipelines, client.TestOptions{})
}
