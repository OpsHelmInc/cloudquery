package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSuggestedResiliencyPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockResiliencehubClient(ctrl)
	var l resiliencehub.ListSuggestedResiliencyPoliciesOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListSuggestedResiliencyPolicies(
		gomock.Any(),
		&resiliencehub.ListSuggestedResiliencyPoliciesInput{},
		gomock.Any(),
	).Return(&l, nil)

	return client.Services{Resiliencehub: mock}
}

func TestResiilencehubSuggestedResiliencyPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, SuggestedResiliencyPolicies(), buildSuggestedResiliencyPolicies, client.TestOptions{})
}
