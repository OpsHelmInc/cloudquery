package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildParameters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)
	var pm types.ParameterMetadata
	require.NoError(t, faker.FakeObject(&pm))

	var tags []types.Tag
	require.NoError(t, faker.FakeObject(&tags))
	mock.EXPECT().DescribeParameters(
		gomock.Any(),
		&ssm.DescribeParametersInput{},
		gomock.Any(),
	).Return(
		&ssm.DescribeParametersOutput{Parameters: []types.ParameterMetadata{pm}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&ssm.ListTagsForResourceInput{
			ResourceId:   pm.Name,
			ResourceType: types.ResourceTypeForTaggingParameter,
		},
		gomock.Any(),
	).Return(
		&ssm.ListTagsForResourceOutput{TagList: tags},
		nil,
	)
	return client.Services{Ssm: mock}
}

func TestParameters(t *testing.T) {
	client.AwsMockTestHelper(t, Parameters(), buildParameters, client.TestOptions{})
}
