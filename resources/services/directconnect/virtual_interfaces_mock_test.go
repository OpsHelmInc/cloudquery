package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDirectconnectVirtualInterfacesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualInterface{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().DescribeVirtualInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualInterfacesOutput{
			VirtualInterfaces: []types.VirtualInterface{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnecVirtualInterfaces(t *testing.T) {
	client.AwsMockTestHelper(t, VirtualInterfaces(), buildDirectconnectVirtualInterfacesMock, client.TestOptions{})
}
