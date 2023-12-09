package glue

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
)

func buildConnections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var connecions glue.GetConnectionsOutput
	if err := faker.FakeObject(&connecions); err != nil {
		t.Fatal(err)
	}
	connecions.NextToken = nil
	m.EXPECT().GetConnections(gomock.Any(), gomock.Any()).Return(&connecions, nil)

	return client.Services{
		Glue: m,
	}
}

func TestConnections(t *testing.T) {
	client.AwsMockTestHelper(t, Connections(), buildConnections, client.TestOptions{})
}
