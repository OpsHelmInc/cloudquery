package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildIotPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	lp := iot.ListPoliciesOutput{}
	require.NoError(t, faker.FakeObject(&lp))

	lp.NextMarker = nil
	m.EXPECT().ListPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p := iot.GetPolicyOutput{}
	require.NoError(t, faker.FakeObject(&p))

	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, Policies(), buildIotPolicies, client.TestOptions{})
}
