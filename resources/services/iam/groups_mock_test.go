package iam

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	err := faker.FakeObject(&g)
	g.Arn = aws.String("arn:aws:iam::123456789012:group/Admins")
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAttachedGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedGroupPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	// list policies
	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupPoliciesOutput{
			PolicyNames: l,
		}, nil)

	// get policy
	gp := iam.GetGroupPolicyOutput{}
	err = faker.FakeObject(&gp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	gp.PolicyDocument = &document
	m.EXPECT().GetGroupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&gp, nil)

	// get group
	user := iamTypes.User{Arn: aws.String("arn:aws:iam::123456789012:user/test")}
	gg := iam.GetGroupOutput{Group: &g, Users: []iamTypes.User{user}}
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&gg, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamGroups(t *testing.T) {
	client.AwsMockTestHelper(t, Groups(), buildIamGroups, client.TestOptions{})
}
