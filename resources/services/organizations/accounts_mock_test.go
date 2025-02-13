package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)

	tt := make([]organizationsTypes.Tag, 3)
	require.NoError(t, faker.FakeObject(&tt))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListTagsForResourceOutput{
			Tags: tt,
		}, nil)

	p := organizationsTypes.Parent{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().ListParents(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&organizations.ListParentsOutput{
			Parents: []organizationsTypes.Parent{p},
		}, nil)
	return client.Services{
		Organizations: m,
	}
}

func TestOrganizationsAccounts(t *testing.T) {
	client.AwsMockTestHelper(t, Accounts(), buildOrganizationsAccounts, client.TestOptions{})
}
