package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildSamplingRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockXrayClient(ctrl)

	test := "test"

	var rule types.SamplingRuleRecord
	require.NoError(t, faker.FakeObject(&rule))

	mock.EXPECT().GetSamplingRules(
		gomock.Any(),
		&xray.GetSamplingRulesInput{},
		gomock.Any(),
	).Return(
		&xray.GetSamplingRulesOutput{
			SamplingRuleRecords: []types.SamplingRuleRecord{
				rule,
			},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&xray.ListTagsForResourceInput{ResourceARN: rule.SamplingRule.RuleARN},
		gomock.Any(),
	).Return(
		&xray.ListTagsForResourceOutput{
			Tags: []types.Tag{
				{
					Key:   &test,
					Value: &test,
				},
			},
		},
		nil,
	)

	return client.Services{Xray: mock}
}

func TestXraySamplingRules(t *testing.T) {
	client.AwsMockTestHelper(t, SamplingRules(), buildSamplingRules, client.TestOptions{})
}
