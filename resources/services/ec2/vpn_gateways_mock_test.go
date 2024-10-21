package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEc2VpnGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.VpnGateway{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().DescribeVpnGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpnGatewaysOutput{
			VpnGateways: []types.VpnGateway{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpnGateways(t *testing.T) {
	client.AwsMockTestHelper(t, VpnGateways(), buildEc2VpnGateways, client.TestOptions{})
}
