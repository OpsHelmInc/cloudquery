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

func buildEc2VpcEndpoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	e := types.VpcEndpoint{}
	require.NoError(t, faker.FakeObject(&e))

	m.EXPECT().DescribeVpcEndpoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointsOutput{
			VpcEndpoints: []types.VpcEndpoint{e},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpoints(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpoints(), buildEc2VpcEndpoints, client.TestOptions{})
}
