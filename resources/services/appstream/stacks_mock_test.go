package appstream

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAppstreamStacksMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAppstreamClient(ctrl)
	stack := types.Stack{}
	require.NoError(t, faker.FakeObject(&stack))

	stackEntitlements := types.Entitlement{}
	require.NoError(t, faker.FakeObject(&stackEntitlements))

	stackUserAssociations := types.UserStackAssociation{}
	require.NoError(t, faker.FakeObject(&stackUserAssociations))

	m.EXPECT().DescribeStacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeStacksOutput{
			Stacks: []types.Stack{stack},
		}, nil)

	m.EXPECT().DescribeEntitlements(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeEntitlementsOutput{
			Entitlements: []types.Entitlement{stackEntitlements},
		}, nil)

	m.EXPECT().DescribeUserStackAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&appstream.DescribeUserStackAssociationsOutput{
			UserStackAssociations: []types.UserStackAssociation{stackUserAssociations},
		}, nil)

	return client.Services{
		Appstream: m,
	}
}

func TestAppstreamStacks(t *testing.T) {
	client.AwsMockTestHelper(t, Stacks(), buildAppstreamStacksMock, client.TestOptions{})
}
