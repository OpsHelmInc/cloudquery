package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildEc2LaunchTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	describeLaunchTemplatesOutput := ec2.DescribeLaunchTemplatesOutput{}
	err := faker.FakeObject(&describeLaunchTemplatesOutput)
	if err != nil {
		t.Fatal(err)
	}
	describeLaunchTemplatesOutput.NextToken = nil
	m.EXPECT().DescribeLaunchTemplates(gomock.Any(), gomock.Any()).Return(&describeLaunchTemplatesOutput, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2LaunchTemplates(t *testing.T) {
	client.AwsMockTestHelper(
		t,
		LaunchTemplates(),
		buildEc2LaunchTemplates,
		client.TestOptions{},
	)
}
