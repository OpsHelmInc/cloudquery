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

func buildInventories(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InventoryResultEntity
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInventory(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.GetInventoryOutput{Entities: []types.InventoryResultEntity{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestInventories(t *testing.T) {
	client.AwsMockTestHelper(t, Inventories(), buildInventories, client.TestOptions{})
}
