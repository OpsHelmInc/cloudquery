package frauddetector

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRules(t *testing.T, client *mocks.MockFrauddetectorClient) {
	data := types.RuleDetail{}
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	client.EXPECT().GetRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.GetRulesOutput{RuleDetails: []types.RuleDetail{data}}, nil,
	)
}
