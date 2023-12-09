package iam

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
)

func buildIamSAMLProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.SAMLProviderListEntry{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSAMLProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListSAMLProvidersOutput{
			SAMLProviderList: []iamTypes.SAMLProviderListEntry{l},
		}, nil)

	p := iam.GetSAMLProviderOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetSAMLProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIAMSamlIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, SamlIdentityProviders(), buildIamSAMLProviders, client.TestOptions{})
}
