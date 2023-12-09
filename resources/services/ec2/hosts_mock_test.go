package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
)

func buildEc2Hosts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	g := types.Host{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

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
