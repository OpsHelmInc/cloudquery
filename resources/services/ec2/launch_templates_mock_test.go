package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
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

	launchTemplate := &ec2.GetLaunchTemplateDataOutput{}
	err = faker.FakeObject(&launchTemplate)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetLaunchTemplateData(gomock.Any(), gomock.Any()).Return(&launchTemplate, nil)

	return client.Services{
		Ec2: m,
	}
}
