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

func buildComplianceSummaryItems(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.ComplianceSummaryItem
	require.NoError(t, faker.FakeObject(&i))

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
