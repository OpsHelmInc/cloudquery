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

func buildSageMakerApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSagemakerClient(ctrl)

	appDets := types.AppDetails{}
	require.NoError(t, faker.FakeObject(&appDets))

	m.EXPECT().ListApps(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sagemaker.ListAppsOutput{Apps: []types.AppDetails{appDets}},
		nil,
	)

	app := sagemaker.DescribeAppOutput{}
	require.NoError(t, faker.FakeObject(&app))

	m.EXPECT().DescribeApp(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&app,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	require.NoError(t, faker.FakeObject(&tagsOut))

	m.EXPECT().ListTags(gomock.Any(), &sagemaker.ListTagsInput{ResourceArn: app.AppArn}, gomock.Any()).Return(
		&tagsOut, nil,
	)

	return client.Services{
		Sagemaker: m,
	}
}

func TestSageMakerApps(t *testing.T) {
	client.AwsMockTestHelper(t, Apps(), buildSageMakerApps, client.TestOptions{})
}
