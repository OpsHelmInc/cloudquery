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

func buildReservedEc2Instances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.ReservedInstances{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeReservedInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeReservedInstancesOutput{
			ReservedInstances: []types.ReservedInstances{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestReservedEc2Instances(t *testing.T) {
	client.AwsMockTestHelper(t, ReservedInstances(), buildReservedEc2Instances, client.TestOptions{})
}
