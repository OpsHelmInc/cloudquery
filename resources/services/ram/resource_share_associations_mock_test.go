package ram

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/golang/mock/gomock"
)

func buildRamResourceShareAssociationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)

	object := types.ResourceShareAssociation{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetResourceShareAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceShareAssociationsOutput{ResourceShareAssociations: []types.ResourceShareAssociation{object}}, nil).MinTimes(1)
	return client.Services{Ram: m}
}

func TestRamResourceShareAssociatedResources(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceShareAssociations(), buildRamResourceShareAssociationsMock, client.TestOptions{})
}
