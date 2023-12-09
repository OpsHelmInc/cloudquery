package account

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
	"github.com/golang/mock/gomock"
)

func buildContacts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAccountClient(ctrl)

	var ci types.ContactInformation
	if err := faker.FakeObject(&ci); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetContactInformation(
		gomock.Any(),
		&account.GetContactInformationInput{},
		gomock.Any(),
	).Return(
		&account.GetContactInformationOutput{ContactInformation: &ci},
		nil,
	)

	return client.Services{Account: mock}
}

func TestContacts(t *testing.T) {
	client.AwsMockTestHelper(t, Contacts(), buildContacts, client.TestOptions{})
}
