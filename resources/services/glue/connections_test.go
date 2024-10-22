package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var connecions glue.GetConnectionsOutput
	require.NoError(t, faker.FakeObject(&connecions))

	connecions.NextToken = nil
	m.EXPECT().GetConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(&connecions, nil)

	return client.Services{
		Glue: m,
	}
}

func TestConnections(t *testing.T) {
	client.AwsMockTestHelper(t, Connections(), buildConnections, client.TestOptions{})
}
