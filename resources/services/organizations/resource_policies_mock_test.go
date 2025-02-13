package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildResourcePolicy(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)

	o := organizations.DescribeResourcePolicyOutput{}
	require.NoError(t, faker.FakeObject(&o))

	m.EXPECT().DescribeResourcePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&o, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestResourcePolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResourcePolicies(), buildResourcePolicy, client.TestOptions{})
}
