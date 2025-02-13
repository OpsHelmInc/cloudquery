package wafregional

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafregionalClient(ctrl)

	var r types.Rule
	require.NoError(t, faker.FakeObject(&r))

	m.EXPECT().ListRules(
		gomock.Any(),
		&wafregional.ListRulesInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRulesOutput{
			Rules: []types.RuleSummary{{RuleId: r.RuleId}},
		},
		nil,
	)

	m.EXPECT().GetRule(
		gomock.Any(),
		&wafregional.GetRuleInput{RuleId: r.RuleId},
		gomock.Any(),
	).Return(
		&wafregional.GetRuleOutput{Rule: &r},
		nil,
	)

	arn := aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:rule/%v", *r.RuleId))

	tif := types.TagInfoForResource{}
	require.NoError(t, faker.FakeObject(&tif))
	tif.ResourceARN = arn
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: arn,
		},
		gomock.Any(),
	).Return(
		&wafregional.ListTagsForResourceOutput{TagInfoForResource: &tif},
		nil,
	)

	return client.Services{Wafregional: m}
}

func TestRules(t *testing.T) {
	client.AwsMockTestHelper(t, Rules(), buildRulesMock, client.TestOptions{})
}
