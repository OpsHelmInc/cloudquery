package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildApprunnerVpcIngressConnectionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockApprunnerClient(ctrl)
	vc := types.VpcIngressConnection{}
	require.NoError(t, faker.FakeObject(&vc))

	m.EXPECT().ListVpcIngressConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListVpcIngressConnectionsOutput{
			VpcIngressConnectionSummaryList: []types.VpcIngressConnectionSummary{{ServiceArn: vc.ServiceArn, VpcIngressConnectionArn: vc.VpcIngressConnectionArn}},
		}, nil)
	m.EXPECT().DescribeVpcIngressConnection(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.DescribeVpcIngressConnectionOutput{VpcIngressConnection: &vc}, nil)

	tags := types.Tag{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&apprunner.ListTagsForResourceOutput{Tags: []types.Tag{tags}}, nil)

	return client.Services{
		Apprunner: m,
	}
}

func TestApprunnerVpcIngressConnector(t *testing.T) {
	client.AwsMockTestHelper(t, VpcIngressConnections(), buildApprunnerVpcIngressConnectionsMock, client.TestOptions{})
}
