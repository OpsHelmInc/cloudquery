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

func buildAppstreamAppBlocksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	object := types.AppBlock{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().DescribeAppBlocks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeAppBlocksOutput{
			AppBlocks: []types.AppBlock{object},
		}, nil)

	tagsOutput := appstream.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tagsOutput))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any()).Return(&tagsOutput, nil).AnyTimes()

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamAppBlocks(t *testing.T) {
	client.AwsMockTestHelper(t, AppBlocks(), buildAppstreamAppBlocksMock, client.TestOptions{})
}
