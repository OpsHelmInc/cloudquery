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

func buildRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	r := iamTypes.Role{}
	require.NoError(t, faker.FakeObject(&r))

	p := iamTypes.AttachedPolicy{}
	require.NoError(t, faker.FakeObject(&p))

	// generate valid json
	document := `{"stuff": 3}`
	r.AssumeRolePolicyDocument = &document

	pd := iam.GetRolePolicyOutput{}
	require.NoError(t, faker.FakeObject(&pd))
	pd.PolicyDocument = &document

	var l []string
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().GetRole(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetRoleOutput{
			Role: &r,
		}, nil)
	m.EXPECT().ListRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListRolePoliciesOutput{
			PolicyNames: l,
		}, nil)
	m.EXPECT().GetRolePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&pd, nil)

	m.EXPECT().ListRoles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListRolesOutput{
			Roles: []iamTypes.Role{r},
		}, nil)
	m.EXPECT().ListAttachedRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedRolePoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	tag := iamTypes.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().ListAttachedRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedRolePoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	// m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.GenerateServiceLastAccessedDetailsOutput{JobId: aws.String("JobId")}, nil)

	// lastAccessed := []iamTypes.ServiceLastAccessed{}
	// require.NoError(t, faker.FakeObject(&lastAccessed))
	// m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
	// 	&iam.GetServiceLastAccessedDetailsOutput{ServicesLastAccessed: lastAccessed, JobStatus: iamTypes.JobStatusTypeCompleted},
	// 	nil,
	// )

	return client.Services{
		Iam: m,
	}
}

func TestIamRoles(t *testing.T) {
	client.AwsMockTestHelper(t, Roles(), buildRoles, client.TestOptions{})
}
