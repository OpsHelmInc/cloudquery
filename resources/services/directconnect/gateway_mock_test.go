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

func buildDirectconnectGatewaysMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.DirectConnectGateway{}
	association := types.DirectConnectGatewayAssociation{}
	attachment := types.DirectConnectGatewayAttachment{}
	require.NoError(t, faker.FakeObject(&l))
	require.NoError(t, faker.FakeObject(&association))
	require.NoError(t, faker.FakeObject(&attachment))
	m.EXPECT().DescribeDirectConnectGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewaysOutput{
			DirectConnectGateways: []types.DirectConnectGateway{l},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAssociationsOutput{
			DirectConnectGatewayAssociations: []types.DirectConnectGatewayAssociation{association},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeDirectConnectGatewayAttachmentsOutput{
			DirectConnectGatewayAttachments: []types.DirectConnectGatewayAttachment{attachment},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectGateways(t *testing.T) {
	client.AwsMockTestHelper(t, Gateways(), buildDirectconnectGatewaysMock, client.TestOptions{})
}
