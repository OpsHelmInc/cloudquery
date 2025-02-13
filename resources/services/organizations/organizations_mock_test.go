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

func buildOrganizations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)

	o := organizations.DescribeOrganizationOutput{}
	require.NoError(t, faker.FakeObject(&o))

	m.EXPECT().DescribeOrganization(gomock.Any(), gomock.Any(), gomock.Any()).Return(&o, nil)

	return client.Services{
		Organizations: m,
	}
}

func TestOrganizations(t *testing.T) {
	client.AwsMockTestHelper(t, Organizations(), buildOrganizations, client.TestOptions{})
}
