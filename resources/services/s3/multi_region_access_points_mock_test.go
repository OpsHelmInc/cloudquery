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

func buildS3MultiRegionAccessPoints(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3controlClient(ctrl)
	mrap := types.MultiRegionAccessPointReport{}
	require.NoError(t, faker.FakeObject(&mrap))

	m.EXPECT().ListMultiRegionAccessPoints(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&s3control.ListMultiRegionAccessPointsOutput{
			AccessPoints: []types.MultiRegionAccessPointReport{mrap},
		}, nil)

	return client.Services{
		S3control: m,
	}
}

func TestMultiRegionAccessPoints(t *testing.T) {
	client.AwsMockTestHelper(t, MultiRegionAccessPoints(), buildS3MultiRegionAccessPoints, client.TestOptions{})
}
