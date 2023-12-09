package kms

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang/mock/gomock"
)

func buildKmsAliases(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockKmsClient(ctrl)

	aliases := kms.ListAliasesOutput{}
	err := faker.FakeObject(&aliases)
	if err != nil {
		t.Fatal(err)
	}
	aliases.NextMarker = nil
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&aliases, nil)

	return client.Services{
		Kms: m,
	}
}

func TestKmsAliases(t *testing.T) {
	client.AwsMockTestHelper(t, Aliases(), buildKmsAliases, client.TestOptions{})
}
