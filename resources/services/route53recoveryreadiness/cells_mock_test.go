package route53recoveryreadiness

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCells(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53recoveryreadinessClient(ctrl)
	co := types.CellOutput{}
	require.NoError(t, faker.FakeObject(&co))

	m.EXPECT().ListCells(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53recoveryreadiness.ListCellsOutput{
			Cells: []types.CellOutput{co},
		}, nil)

	return client.Services{
		Route53recoveryreadiness: m,
	}
}

func TestCells(t *testing.T) {
	client.AwsMockTestHelper(t, Cells(), buildCells, client.TestOptions{Region: "us-west-2"})
}
