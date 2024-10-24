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

func buildRateBasedRulesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafregionalClient(ctrl)

	var rule types.RateBasedRule
	require.NoError(t, faker.FakeObject(&rule))

	m.EXPECT().ListRateBasedRules(
		gomock.Any(),
		&wafregional.ListRateBasedRulesInput{},
		gomock.Any(),
	).Return(
		&wafregional.ListRateBasedRulesOutput{
			Rules: []types.RuleSummary{{RuleId: rule.RuleId}},
		},
		nil,
	)

	m.EXPECT().GetRateBasedRule(
		gomock.Any(),
		&wafregional.GetRateBasedRuleInput{RuleId: rule.RuleId},
		gomock.Any(),
	).Return(
		&wafregional.GetRateBasedRuleOutput{Rule: &rule},
		nil,
	)

	arn := aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:ratebasedrule/%v", *rule.RuleId))

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

func TestRateBasedRules(t *testing.T) {
	client.AwsMockTestHelper(t, RateBasedRules(), buildRateBasedRulesMock, client.TestOptions{})
}
