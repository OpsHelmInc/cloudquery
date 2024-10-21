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

func buildResourceSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoveryreadinessClient(ctrl)
	rs := types.ResourceSetOutput{}
	require.NoError(t, faker.FakeObject(&rs))

	m.EXPECT().ListResourceSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53recoveryreadiness.ListResourceSetsOutput{
			ResourceSets: []types.ResourceSetOutput{rs},
		}, nil)

	return client.Services{
		Route53recoveryreadiness: m,
	}
}

func TestResourceSets(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceSets(), buildResourceSets, client.TestOptions{Region: "us-west-2"})
}
