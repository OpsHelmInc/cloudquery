package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIamGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	require.NoError(t, faker.FakeObject(&g))

	p := iamTypes.AttachedPolicy{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAttachedGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedGroupPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil).MinTimes(1)

	var l []string
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().ListGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupPoliciesOutput{
			PolicyNames: l,
		}, nil)

	gp := iam.GetGroupPolicyOutput{}
	require.NoError(t, faker.FakeObject(&gp))
	document := "{\"test\": {\"t1\":1}}"
	gp.PolicyDocument = &document
	m.EXPECT().GetGroupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&gp, nil)

	// m.EXPECT().GenerateServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.GenerateServiceLastAccessedDetailsOutput{JobId: aws.String("JobId")}, nil)

	// lastAccessed := []iamTypes.ServiceLastAccessed{}
	// require.NoError(t, faker.FakeObject(&lastAccessed))
	// m.EXPECT().GetServiceLastAccessedDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
	// 	&iam.GetServiceLastAccessedDetailsOutput{ServicesLastAccessed: lastAccessed, JobStatus: iamTypes.JobStatusTypeCompleted},
	// 	nil,
	// )

	m.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetGroupOutput{
			Group: &g,
			Users: []iamTypes.User{
				{
					Arn: aws.String("arn:aws:iam::123456789012:user/name"),
				},
			},
		},
		nil,
	)

	return client.Services{
		Iam: m,
	}
}

func TestIamGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildIamGroups, client.TestOptions{})
}
