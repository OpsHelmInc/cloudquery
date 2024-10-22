package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEgressOnlyInternetGateways(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	egressOutput := ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	require.NoError(t, faker.FakeObject(&egressOutput))

	egressOutput.NextToken = nil
	m.EXPECT().DescribeEgressOnlyInternetGateways(gomock.Any(), gomock.Any(), gomock.Any()).Return(&egressOutput, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEgressOnlyInternetGateways(t *testing.T) {
	client.AwsMockTestHelper(t, EgressOnlyInternetGateways(), buildEgressOnlyInternetGateways, client.TestOptions{})
}
