package elastictranscoder

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildElastictranscoderPipelinesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElastictranscoderClient(ctrl)

	pipeline := types.Pipeline{}
	require.NoError(t, faker.FakeObject(&pipeline))

	m.EXPECT().ListPipelines(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elastictranscoder.ListPipelinesOutput{Pipelines: []types.Pipeline{pipeline}},
		nil,
	)

	job := types.Job{}
	require.NoError(t, faker.FakeObject(&job))

	job.PipelineId = pipeline.Id

	m.EXPECT().ListJobsByPipeline(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elastictranscoder.ListJobsByPipelineOutput{Jobs: []types.Job{job}},
		nil,
	)

	return client.Services{
		Elastictranscoder: m,
	}
}

func TestElastictranscoderPipelines(t *testing.T) {
	client.AwsMockTestHelper(t, Pipelines(), buildElastictranscoderPipelinesMock, client.TestOptions{})
}
