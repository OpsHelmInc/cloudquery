package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAnalysesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var lo quicksight.ListAnalysesOutput
	require.NoError(t, faker.FakeObject(&lo))

	lo.NextToken = nil
	m.EXPECT().ListAnalyses(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)

	var qs quicksight.DescribeAnalysisOutput
	require.NoError(t, faker.FakeObject(&qs))

	m.EXPECT().DescribeAnalysis(gomock.Any(), gomock.Any(), gomock.Any()).Return(&qs, nil)

	var to quicksight.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&to))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		Quicksight: m,
	}
}

func TestQuicksightAnalyses(t *testing.T) {
	client.AwsMockTestHelper(t, Analyses(), buildAnalysesMock, client.TestOptions{})
}
