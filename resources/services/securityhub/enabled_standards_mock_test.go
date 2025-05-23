package securityhub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEnabledStandards(t *testing.T, ctrl *gomock.Controller) client.Services {
	shMock := mocks.NewMockSecurityhubClient(ctrl)
	standardsSubscription := types.StandardsSubscription{}
	require.NoError(t, faker.FakeObject(&standardsSubscription))

	shMock.EXPECT().GetEnabledStandards(
		gomock.Any(),
		&securityhub.GetEnabledStandardsInput{MaxResults: aws.Int32(100)},
		gomock.Any(),
	).Return(
		&securityhub.GetEnabledStandardsOutput{StandardsSubscriptions: []types.StandardsSubscription{standardsSubscription}},
		nil,
	)

	return client.Services{Securityhub: shMock}
}

func TestEnabledStandards(t *testing.T) {
	client.AwsMockTestHelper(t, EnabledStandards(), buildEnabledStandards, client.TestOptions{})
}
