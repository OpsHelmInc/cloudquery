package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildStudios(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.StudioSummary
	require.NoError(t, faker.FakeObject(&summary))

	mock.EXPECT().ListStudios(gomock.Any(), &emr.ListStudiosInput{}, gomock.Any()).Return(
		&emr.ListStudiosOutput{Studios: []types.StudioSummary{summary}},
		nil,
	)

	var studio types.Studio
	require.NoError(t, faker.FakeObject(&studio))
	mock.EXPECT().DescribeStudio(gomock.Any(), &emr.DescribeStudioInput{StudioId: summary.StudioId}, gomock.Any()).Return(
		&emr.DescribeStudioOutput{Studio: &studio},
		nil,
	)

	var sessionMappingSummary types.SessionMappingSummary
	require.NoError(t, faker.FakeObject(&sessionMappingSummary))
	mock.EXPECT().ListStudioSessionMappings(gomock.Any(), &emr.ListStudioSessionMappingsInput{StudioId: summary.StudioId}, gomock.Any()).Return(
		&emr.ListStudioSessionMappingsOutput{SessionMappings: []types.SessionMappingSummary{sessionMappingSummary}},
		nil,
	)

	var sessionMappingDetail types.SessionMappingDetail
	require.NoError(t, faker.FakeObject(&sessionMappingDetail))
	mock.EXPECT().GetStudioSessionMapping(gomock.Any(), &emr.GetStudioSessionMappingInput{
		StudioId:     sessionMappingSummary.StudioId,
		IdentityType: sessionMappingSummary.IdentityType,
		IdentityId:   sessionMappingSummary.IdentityId,
	}, gomock.Any()).Return(
		&emr.GetStudioSessionMappingOutput{
			SessionMapping: &sessionMappingDetail,
		},
		nil,
	)

	return client.Services{Emr: mock}
}

func TestStudios(t *testing.T) {
	client.AwsMockTestHelper(t, Studios(), buildStudios, client.TestOptions{})
}
