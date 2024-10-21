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

func buildAssociations(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.Association
	require.NoError(t, faker.FakeObject(&i))

	mock.EXPECT().ListAssociations(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.ListAssociationsOutput{Associations: []types.Association{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestAssociations(t *testing.T) {
	client.AwsMockTestHelper(t, Associations(), buildAssociations, client.TestOptions{})
}
