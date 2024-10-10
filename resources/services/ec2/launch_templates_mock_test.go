package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2LaunchTemplates(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	lt := types.LaunchTemplate{}
	require.NoError(t, faker.FakeObject(&lt))

	m.EXPECT().DescribeLaunchTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeLaunchTemplatesOutput{
			LaunchTemplates: []types.LaunchTemplate{lt},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2LaunchTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, LaunchTemplates(), buildEc2LaunchTemplates, client.TestOptions{})
}
