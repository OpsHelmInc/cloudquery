package s3

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildS3AccessGrants(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3controlClient(ctrl)
	ag := types.ListAccessGrantEntry{}
	require.NoError(t, faker.FakeObject(&ag))

	m.EXPECT().ListAccessGrants(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3control.ListAccessGrantsOutput{
			AccessGrantsList: []types.ListAccessGrantEntry{ag},
		}, nil)

	return client.Services{
		S3control: m,
	}
}

func TestAccessGrants(t *testing.T) {
	client.AwsMockTestHelper(t, AccessGrants(), buildS3AccessGrants, client.TestOptions{})
}
