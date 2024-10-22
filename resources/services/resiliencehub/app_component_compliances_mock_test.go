package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAppComponentCompliances(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppComponentCompliancesOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListAppComponentCompliances(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
}
