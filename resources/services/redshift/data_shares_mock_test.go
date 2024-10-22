package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDataSharesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRedshiftClient(ctrl)
	ds := types.DataShare{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().DescribeDataShares(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&redshift.DescribeDataSharesOutput{
			DataShares: []types.DataShare{ds},
		}, nil)

	return client.Services{
		Redshift: m,
	}
}

func TestRedshiftDataShares(t *testing.T) {
	client.AwsMockTestHelper(t, DataShares(), buildDataSharesMock, client.TestOptions{})
}
