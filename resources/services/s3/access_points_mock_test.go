package s3

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildS3AccessPoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3controlClient(ctrl)
	ap := types.AccessPoint{}
	require.NoError(t, faker.FakeObject(&ap))

	m.EXPECT().ListAccessPoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3control.ListAccessPointsOutput{
			AccessPointList: []types.AccessPoint{ap},
		}, nil)

	return client.Services{
		S3control: m,
	}
}

func TestAccessPoints(t *testing.T) {
	client.AwsMockTestHelper(t, AccessPoints(), buildS3AccessPoints, client.TestOptions{})
}
