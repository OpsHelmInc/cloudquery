package iot

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/golang/mock/gomock"
)

func buildIotThingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	thing := types.ThingAttribute{}
	err := faker.FakeObject(&thing)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListThings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iot.ListThingsOutput{Things: []types.ThingAttribute{thing}}, nil)

	lp := iot.ListThingPrincipalsOutput{}
	err = faker.FakeObject(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextToken = nil
	m.EXPECT().ListThingPrincipals(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotThings(t *testing.T) {
	client.AwsMockTestHelper(t, Things(), buildIotThingsMock, client.TestOptions{})
}
