package inspector2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/golang/mock/gomock"
)

func buildInspectorV2Findings(t *testing.T, ctrl *gomock.Controller) client.Services {
	inspectorClient := mocks.NewMockInspector2Client(ctrl)

	finding := types.Finding{}
	err := faker.FakeObject(&finding)
	if err != nil {
		t.Fatal(err)
	}

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&inspector2.ListFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return client.Services{
		Inspector2: inspectorClient,
	}
}

func TestInspectorV2Findings(t *testing.T) {
	client.AwsMockTestHelper(t, Findings(), buildInspectorV2Findings, client.TestOptions{})
}
