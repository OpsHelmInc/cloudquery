package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSageMakerEndpointConfigs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	summ := types.EndpointConfigSummary{}
	require.NoError(t, faker.FakeObject(&summ))

	m.EXPECT().ListEndpointConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListEndpointConfigsOutput{EndpointConfigs: []types.EndpointConfigSummary{summ}},
		nil,
	)

	endpointConfig := sagemaker.DescribeEndpointConfigOutput{}
	require.NoError(t, faker.FakeObject(&endpointConfig))

	m.EXPECT().DescribeEndpointConfig(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&endpointConfig,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	require.NoError(t, faker.FakeObject(&tagsOut))

	tagsOut.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Sagemaker: m,
	}
}

func TestSageMakerEndpointConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, EndpointConfigurations(), buildSageMakerEndpointConfigs, client.TestOptions{})
}
