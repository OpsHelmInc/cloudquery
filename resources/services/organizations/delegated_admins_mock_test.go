package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDelegatedAdministrators(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	da := types.DelegatedAdministrator{}
	require.NoError(t, faker.FakeObject(&da))

	m.EXPECT().ListDelegatedAdministrators(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListDelegatedAdministratorsOutput{
			DelegatedAdministrators: []types.DelegatedAdministrator{da},
		}, nil)

	ds := types.DelegatedService{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().ListDelegatedServicesForAccount(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListDelegatedServicesForAccountOutput{
			DelegatedServices: []types.DelegatedService{ds},
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestDelegatedAdministrators(t *testing.T) {
	client.AwsMockTestHelper(t, DelegatedAdministrators(), buildDelegatedAdministrators, client.TestOptions{})
}
