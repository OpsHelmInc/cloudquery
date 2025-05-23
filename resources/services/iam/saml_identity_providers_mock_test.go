package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIamSAMLProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.SAMLProviderListEntry{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().ListSAMLProviders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListSAMLProvidersOutput{
			SAMLProviderList: []iamTypes.SAMLProviderListEntry{l},
		}, nil)

	p := iam.GetSAMLProviderOutput{}
	require.NoError(t, faker.FakeObject(&p))
	m.EXPECT().GetSAMLProvider(gomock.Any(), gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIAMSamlIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, SamlIdentityProviders(), buildIamSAMLProviders, client.TestOptions{})
}
