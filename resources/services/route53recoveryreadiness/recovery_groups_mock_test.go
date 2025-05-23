package route53recoveryreadiness

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRecoveryGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoveryreadinessClient(ctrl)
	rco := types.RecoveryGroupOutput{}
	require.NoError(t, faker.FakeObject(&rco))

	m.EXPECT().ListRecoveryGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53recoveryreadiness.ListRecoveryGroupsOutput{
			RecoveryGroups: []types.RecoveryGroupOutput{rco},
		}, nil)

	return client.Services{
		Route53recoveryreadiness: m,
	}
}

func TestRecoveryGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RecoveryGroups(), buildRecoveryGroups, client.TestOptions{Region: "us-west-2"})
}
