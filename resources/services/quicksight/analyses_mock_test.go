package quicksight

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/golang/mock/gomock"
)

func buildAnalysesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var lo quicksight.ListAnalysesOutput
	if err := faker.FakeObject(&lo); err != nil {
		t.Fatal(err)
	}
	lo.NextToken = nil
	m.EXPECT().ListAnalyses(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)

	var qs quicksight.DescribeAnalysisOutput
	if err := faker.FakeObject(&qs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeAnalysis(gomock.Any(), gomock.Any(), gomock.Any()).Return(&qs, nil)

	var to quicksight.ListTagsForResourceOutput
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightAnalyses(t *testing.T) {
	client.AwsMockTestHelper(t, Analyses(), buildAnalysesMock, client.TestOptions{})
}
