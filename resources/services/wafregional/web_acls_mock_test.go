package wafregional

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildWebACLsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafregionalClient(ctrl)

	var acl types.WebACL
	require.NoError(t, faker.FakeObject(&acl))

	m.EXPECT().ListWebACLs(
		gomock.Any(),
		&wafregional.ListWebACLsInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{{WebACLId: acl.WebACLId}},
		},
		nil,
	)

	m.EXPECT().GetWebACL(
		gomock.Any(),
		&wafregional.GetWebACLInput{WebACLId: acl.WebACLId},
		gomock.Any(),
	).Return(
		&wafregional.GetWebACLOutput{WebACL: &acl},
		nil,
	)

	tif := types.TagInfoForResource{}
	require.NoError(t, faker.FakeObject(&tif))
	tif.ResourceARN = acl.WebACLArn

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: acl.WebACLArn,
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{TagInfoForResource: &tif},
		nil,
	)

	m.EXPECT().ListResourcesForWebACL(
		gomock.Any(),
		&wafregional.ListResourcesForWebACLInput{
			WebACLId: acl.WebACLId,
		},
		gomock.Any(),
	).Return(
		&wafregional.ListResourcesForWebACLOutput{
			ResourceArns: []string{"arn:aws:cloudfront::123456789012:distribution/EDFDVBD6EXAMPLE"},
		},
		nil,
	)

	return client.Services{Wafregional: m}
}

func TestWebACLs(t *testing.T) {
	client.AwsMockTestHelper(t, WebAcls(), buildWebACLsMock, client.TestOptions{})
}
