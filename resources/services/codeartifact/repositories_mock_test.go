package codeartifact

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRepositories(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCodeartifactClient(ctrl)

	repoSummary := types.RepositorySummary{}
	require.NoError(t, faker.FakeObject(&repoSummary))

	repo := types.RepositoryDescription{}
	require.NoError(t, faker.FakeObject(&repo))

	m.EXPECT().ListRepositories(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListRepositoriesOutput{
			Repositories: []types.RepositorySummary{repoSummary},
			NextToken:    nil,
		},
		nil,
	)

	m.EXPECT().DescribeRepository(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.DescribeRepositoryOutput{
			Repository: &repo,
		},
		nil,
	)

	tag := types.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&codeartifact.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		},
		nil,
	)

	return client.Services{Codeartifact: m}
}

func TestRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildRepositories, client.TestOptions{})
}
