package ram

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildRamResourceSharesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	object := types.ResourceShare{}
	require.NoError(t, faker.FakeObject(&object))

	m.EXPECT().GetResourceShares(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceSharesOutput{ResourceShares: []types.ResourceShare{object}}, nil).MinTimes(1)

	summary := types.ResourceSharePermissionSummary{}
	require.NoError(t, faker.FakeObject(&summary))

	var version int32
	require.NoError(t, faker.FakeObject(&version))

	verStr := fmt.Sprint(version)
	summary.Version = &verStr

	m.EXPECT().ListResourceSharePermissions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.ListResourceSharePermissionsOutput{Permissions: []types.ResourceSharePermissionSummary{summary}}, nil).MinTimes(1)

	detail := `{"key":"value"}`

	m.EXPECT().GetPermission(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetPermissionOutput{Permission: &types.ResourceSharePermissionDetail{Permission: &detail}}, nil).MinTimes(1)

	return client.Services{Ram: m}
}

func TestRamResourceShares(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceShares(), buildRamResourceSharesMock, client.TestOptions{})
}
