package resiliencehub

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildComponentRecommendations(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppComponentRecommendationsOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListAppComponentRecommendations(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
}
