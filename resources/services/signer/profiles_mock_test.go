package signer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildProfiles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSignerClient(ctrl)
	profileList := types.SigningProfile{}
	require.NoError(t, faker.FakeObject(&profileList))
	m.EXPECT().ListSigningProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&signer.ListSigningProfilesOutput{
			Profiles: []types.SigningProfile{profileList},
		}, nil)

	profile := signer.GetSigningProfileOutput{}
	require.NoError(t, faker.FakeObject(&profile))
	m.EXPECT().GetSigningProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&profile, nil)

	return client.Services{
		Signer: m,
	}
}

func TestProfiles(t *testing.T) {
	client.AwsMockTestHelper(t, Profiles(), buildProfiles, client.TestOptions{})
}
