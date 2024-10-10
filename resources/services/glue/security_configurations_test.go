package glue

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSecurityConfigurationsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var s glue.GetSecurityConfigurationsOutput
	require.NoError(t, faker.FakeObject(&s))
	s.NextToken = nil
	m.EXPECT().GetSecurityConfigurations(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&s, nil)

	return client.Services{
		Glue: m,
	}
}

func TestSecurityConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, SecurityConfigurations(), buildSecurityConfigurationsMock, client.TestOptions{})
}
