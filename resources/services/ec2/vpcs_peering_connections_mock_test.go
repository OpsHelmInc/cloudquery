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

func buildEc2VpcsPeeringConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.VpcPeeringConnection{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcPeeringConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcPeeringConnectionsOutput{
			VpcPeeringConnections: []types.VpcPeeringConnection{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcsPeeringConnections(t *testing.T) {
	client.AwsMockTestHelper(t, VpcPeeringConnections(), buildEc2VpcsPeeringConnections, client.TestOptions{})
}
