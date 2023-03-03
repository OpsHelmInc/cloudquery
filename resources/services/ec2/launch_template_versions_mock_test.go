package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2LaunchTemplateVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	describeLaunchTemplateVersionsOutput := ec2.DescribeLaunchTemplateVersionsOutput{}
	err := faker.FakeObject(&describeLaunchTemplateVersionsOutput)
	if err != nil {
		t.Fatal(err)
	}
	describeLaunchTemplateVersionsOutput.NextToken = nil
	m.EXPECT().DescribeLaunchTemplateVersions(gomock.Any(), gomock.Any()).Return(&describeLaunchTemplateVersionsOutput, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2LaunchTemplateVersions(t *testing.T) {
	client.AwsMockTestHelper(
		t,
		LaunchTemplateVersions(),
		buildEc2LaunchTemplateVersions,
		client.TestOptions{},
	)
}
