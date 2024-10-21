package auditmanager

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildAssessments(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockAuditmanagerClient(ctrl)
	ami := types.AssessmentMetadataItem{}
	require.NoError(t, faker.FakeObject(&ami))

	m.EXPECT().ListAssessments(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&auditmanager.ListAssessmentsOutput{
			AssessmentMetadata: []types.AssessmentMetadataItem{ami},
		}, nil)

	assessment := types.Assessment{}
	require.NoError(t, faker.FakeObject(&assessment))

	m.EXPECT().GetAssessment(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&auditmanager.GetAssessmentOutput{
			Assessment: &assessment,
		}, nil)
	return client.Services{
		Auditmanager: m,
	}
}

func TestAssessments(t *testing.T) {
	client.AwsMockTestHelper(t, Assessments(), buildAssessments, client.TestOptions{})
}
