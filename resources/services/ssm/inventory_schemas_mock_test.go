package ssm

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildInventorySchemas(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InventoryItemSchema
	require.NoError(t, faker.FakeObject(&i))

	mock.EXPECT().GetInventorySchema(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.GetInventorySchemaOutput{Schemas: []types.InventoryItemSchema{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestInventorySchemas(t *testing.T) {
	client.AwsMockTestHelper(t, InventorySchemas(), buildInventorySchemas, client.TestOptions{})
}
