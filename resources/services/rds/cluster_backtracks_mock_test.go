package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRdsClusterBacktracks(t *testing.T, mockRds *mocks.MockRdsClient) {
	var d types.DBClusterBacktrack
	require.NoError(t, faker.FakeObject(&d))

	mockRds.EXPECT().DescribeDBClusterBacktracks(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rds.DescribeDBClusterBacktracksOutput{
		DBClusterBacktracks: []types.DBClusterBacktrack{d},
	}, nil)
}
