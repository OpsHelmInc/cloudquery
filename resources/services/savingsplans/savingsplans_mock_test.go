package savingsplans

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSavingPlansPlans(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSavingsplansClient(ctrl)

	var s types.SavingsPlan
	require.NoError(t, faker.FakeObject(&s))

	m.EXPECT().DescribeSavingsPlans(
		gomock.Any(),
		&savingsplans.DescribeSavingsPlansInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).Return(
		&savingsplans.DescribeSavingsPlansOutput{
			SavingsPlans: []types.SavingsPlan{s},
		},
		nil,
	)

	return client.Services{
		Savingsplans: m,
	}
}

func TestSavingsplansPlans(t *testing.T) {
	client.AwsMockTestHelper(t, Plans(), buildSavingPlansPlans, client.TestOptions{})
}
