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

func buildS3AccessGrantInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3controlClient(ctrl)
	ag := types.ListAccessGrantsInstanceEntry{}
	require.NoError(t, faker.FakeObject(&ag))

	m.EXPECT().ListAccessGrantsInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3control.ListAccessGrantsInstancesOutput{
			AccessGrantsInstancesList: []types.ListAccessGrantsInstanceEntry{ag},
		}, nil)

	return client.Services{
		S3control: m,
	}
}

func TestAccessGrantInstances(t *testing.T) {
	client.AwsMockTestHelper(t, AccessGrantInstances(), buildS3AccessGrantInstances, client.TestOptions{})
}
