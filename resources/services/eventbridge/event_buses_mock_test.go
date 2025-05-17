package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildEventBridgeEventBusesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEventbridgeClient(ctrl)
	bus := types.EventBus{}
	err := faker.FakeObject(&bus)
	if err != nil {
		t.Fatal(err)
	}

	busDescription := eventbridge.DescribeEventBusOutput{}
	require.NoError(t, faker.FakeObject(&busDescription))

	rule := types.Rule{}
	err = faker.FakeObject(&rule)
	if err != nil {
		t.Fatal(err)
	}

	ruleDescription := eventbridge.DescribeRuleOutput{}
	require.NoError(t, faker.FakeObject(&ruleDescription))

	tags := eventbridge.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEventBuses(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListEventBusesOutput{
			EventBuses: []types.EventBus{bus},
		}, nil)

	m.EXPECT().DescribeEventBus(gomock.Any(), &eventbridge.DescribeEventBusInput{Name: bus.Name}).Return(&busDescription, nil)
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&eventbridge.ListRulesOutput{
			Rules: []types.Rule{rule},
		}, nil)
	m.EXPECT().DescribeRule(gomock.Any(), &eventbridge.DescribeRuleInput{Name: rule.Name, EventBusName: busDescription.Name}).Return(&ruleDescription, nil)

	// bus tags
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(2).Return(
		&tags, nil)

	return client.Services{
		Eventbridge: m,
	}
}

func TestEventBridgeEventBuses(t *testing.T) {
	client.AwsMockTestHelper(t, EventBuses(), buildEventBridgeEventBusesMock, client.TestOptions{})
}
