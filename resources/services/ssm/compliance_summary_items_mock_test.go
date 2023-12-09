package ssm

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
)

func buildComplianceSummaryItems(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.ComplianceSummaryItem
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListComplianceSummaries(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.ListComplianceSummariesOutput{ComplianceSummaryItems: []types.ComplianceSummaryItem{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestComplianceSummaryItems(t *testing.T) {
	client.AwsMockTestHelper(t, ComplianceSummaryItems(), buildComplianceSummaryItems, client.TestOptions{})
}
