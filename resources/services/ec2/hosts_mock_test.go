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

func buildEc2Hosts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := types.Host{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().DescribeHosts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeHostsOutput{
			Hosts: []types.Host{g},
		}, nil)

	services := client.Services{
		Ec2: m,
	}
	return services
}

func TestEc2Hosts(t *testing.T) {
	client.AwsMockTestHelper(t, Hosts(), buildEc2Hosts, client.TestOptions{})
}
