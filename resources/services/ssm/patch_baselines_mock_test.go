package ssm

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
)

func buildPatchBaselines(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.PatchBaselineIdentity
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribePatchBaselines(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&ssm.DescribePatchBaselinesOutput{BaselineIdentities: []types.PatchBaselineIdentity{i}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestPatchBaselines(t *testing.T) {
	client.AwsMockTestHelper(t, PatchBaselines(), buildPatchBaselines, client.TestOptions{})
}
