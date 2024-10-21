package batch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildBatchJobsMock(t *testing.T, m *mocks.MockBatchClient) client.Services {
	services := client.Services{
		Batch: m,
	}
	a := types.JobSummary{}
	require.NoError(t, faker.FakeObject(&a))

	d := types.JobDetail{}
	require.NoError(t, faker.FakeObject(&d))

	m.EXPECT().ListJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&batch.ListJobsOutput{
			JobSummaryList: []types.JobSummary{a},
		}, nil).Times(len(allJobStatuses))

	m.EXPECT().DescribeJobs(gomock.Any(), &batch.DescribeJobsInput{
		Jobs: []string{*a.JobId},
	}, gomock.Any()).Return(
		&batch.DescribeJobsOutput{
			Jobs: []types.JobDetail{d},
		}, nil).Times(len(allJobStatuses))

	tagResponse := batch.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil).Times(len(allJobStatuses))

	return services
}
