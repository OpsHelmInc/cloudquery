package shield

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/golang/mock/gomock"
)

func buildAttacks(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	protection := shield.ListAttacksOutput{}
	err := faker.FakeObject(&protection)
	if err != nil {
		t.Fatal(err)
	}
	protection.NextToken = nil
	m.EXPECT().ListAttacks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&protection, nil)

	tags := shield.DescribeAttackOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeAttack(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)
	return client.Services{
		Shield: m,
	}
}

func TestAttacks(t *testing.T) {
	client.AwsMockTestHelper(t, Attacks(), buildAttacks, client.TestOptions{})
}
