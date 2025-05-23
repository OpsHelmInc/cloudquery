package ecrpublic

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEcrPublicRepositoriesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrpublicClient(ctrl)
	l := types.Repository{}
	require.NoError(t, faker.FakeObject(&l))
	i := types.ImageDetail{}
	require.NoError(t, faker.FakeObject(&i))

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecrpublic.DescribeRepositoriesOutput{
			Repositories: []types.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecrpublic.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	tagResponse := ecrpublic.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagResponse))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tagResponse, nil)

	return client.Services{
		Ecrpublic: m,
	}
}

func TestEcrPublicRepositories(t *testing.T) {
	client.AwsMockTestHelper(t, Repositories(), buildEcrPublicRepositoriesMock, client.TestOptions{})
}
