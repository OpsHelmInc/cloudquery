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

func buildEc2NetworkInterfaces(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	niOutput := ec2.DescribeNetworkInterfacesOutput{}
	require.NoError(t, faker.FakeObject(&niOutput))

	niOutput.NextToken = nil
	m.EXPECT().DescribeNetworkInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(&niOutput, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2NetworkInterfaces(t *testing.T) {
	client.AwsMockTestHelper(t, NetworkInterfaces(), buildEc2NetworkInterfaces, client.TestOptions{})
}
