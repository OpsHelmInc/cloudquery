package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEc2NetworkAcls(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	l := types.NetworkAcl{}
	require.NoError(t, faker.FakeObject(&l))

	l.IsDefault = aws.Bool(false)
	m.EXPECT().DescribeNetworkAcls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNetworkAclsOutput{
			NetworkAcls: []types.NetworkAcl{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2NetworkAclsMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, NetworkAcls(), buildEc2NetworkAcls, client.TestOptions{})
}
