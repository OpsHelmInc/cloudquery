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

func buildComplianceDetails(t *testing.T, m *mocks.MockConfigserviceClient) client.Services {
	l := types.EvaluationResult{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().GetComplianceDetailsByConfigRule(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.GetComplianceDetailsByConfigRuleOutput{
			EvaluationResults: []types.EvaluationResult{l},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}
