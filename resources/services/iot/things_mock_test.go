package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIotThingsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	thing := types.ThingAttribute{}
	require.NoError(t, faker.FakeObject(&thing))
	m.EXPECT().ListThings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iot.ListThingsOutput{Things: []types.ThingAttribute{thing}}, nil)

	lp := iot.ListThingPrincipalsOutput{}
	require.NoError(t, faker.FakeObject(&lp))
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
