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

func buildEc2VpcEndpointServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	sd := types.ServiceDetail{}
	require.NoError(t, faker.FakeObject(&sd))

	m.EXPECT().DescribeVpcEndpointServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServicesOutput{
			ServiceDetails: []types.ServiceDetail{sd},
		}, nil)

	ap := types.AllowedPrincipal{}
	require.NoError(t, faker.FakeObject(&ap))

	m.EXPECT().DescribeVpcEndpointServicePermissions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeVpcEndpointServicePermissionsOutput{
			AllowedPrincipals: []types.AllowedPrincipal{ap},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2VpcEndpointServices(t *testing.T) {
	client.AwsMockTestHelper(t, VpcEndpointServices(), buildEc2VpcEndpointServices, client.TestOptions{})
}
