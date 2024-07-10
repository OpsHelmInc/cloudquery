package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2LaunchTemplateVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	ltv := types.LaunchTemplateVersion{}
	require.NoError(t, faker.FakeObject(&ltv))

	m.EXPECT().DescribeLaunchTemplateVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeLaunchTemplateVersionsOutput{
			LaunchTemplateVersions: []types.LaunchTemplateVersion{ltv},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2LaunchTemplateVersions(t *testing.T) {
	client.AwsMockTestHelper(t, LaunchTemplateVersions(), buildEc2LaunchTemplateVersions, client.TestOptions{})
}