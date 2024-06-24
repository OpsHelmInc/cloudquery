package account

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAlternateContacts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAccountClient(ctrl)

	var ac types.AlternateContact
	require.NoError(t, faker.FakeObject(&ac))

	mock.EXPECT().GetAlternateContact(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&account.GetAlternateContactOutput{AlternateContact: &ac},
		nil,
	).MinTimes(1)

	return client.Services{Account: mock}
}

func TestAlternateContacts(t *testing.T) {
	client.AwsMockTestHelper(t, AlternateContacts(), buildAlternateContacts, client.TestOptions{})
}
