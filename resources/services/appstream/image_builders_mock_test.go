package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAppstreamImageBuildersMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.ImageBuilder{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeImageBuilders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeImageBuildersOutput{
			ImageBuilders: []types.ImageBuilder{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamImageBuilders(t *testing.T) {
	client.AwsMockTestHelper(t, ImageBuilders(), buildAppstreamImageBuildersMock, client.TestOptions{})
}
