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

func buildIamInstanceProfiles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	p := iamTypes.InstanceProfile{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().ListInstanceProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListInstanceProfilesOutput{
			InstanceProfiles: []iamTypes.InstanceProfile{p},
		}, nil)

	tag := iamTypes.Tag{}
	require.NoError(t, faker.FakeObject(&tag))
	m.EXPECT().ListInstanceProfileTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListInstanceProfileTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamInstanceProfiles(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceProfiles(), buildIamInstanceProfiles, client.TestOptions{})
}
