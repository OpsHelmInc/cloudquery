package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildConfigConformancePack(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)

	var cpd types.ConformancePackDetail
	require.NoError(t, faker.FakeObject(&cpd))

	var cprc types.ConformancePackRuleCompliance
	require.NoError(t, faker.FakeObject(&cprc))

	var cpre types.ConformancePackEvaluationResult
	require.NoError(t, faker.FakeObject(&cpre))

	m.EXPECT().DescribeConformancePacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConformancePacksOutput{
			ConformancePackDetails: []types.ConformancePackDetail{cpd},
		}, nil)
	m.EXPECT().DescribeConformancePackCompliance(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConformancePackComplianceOutput{
			ConformancePackRuleComplianceList: []types.ConformancePackRuleCompliance{cprc},
		}, nil)
	m.EXPECT().GetConformancePackComplianceDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.GetConformancePackComplianceDetailsOutput{
			ConformancePackRuleEvaluationResults: []types.ConformancePackEvaluationResult{cpre},
		}, nil)

	return client.Services{
		Configservice: m,
	}
}

func TestConfigConformancePack(t *testing.T) {
	client.AwsMockTestHelper(t, ConformancePacks(), buildConfigConformancePack, client.TestOptions{})
}
