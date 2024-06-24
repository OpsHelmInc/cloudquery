package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDynamodbGlobalTablesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var globalTable types.GlobalTable
	require.NoError(t, faker.FakeObject(&globalTable))

	listOutput := &dynamodb.ListGlobalTablesOutput{
		GlobalTables: []types.GlobalTable{globalTable},
	}
	m.EXPECT().ListGlobalTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)
	var gtd types.GlobalTableDescription
	require.NoError(t, faker.FakeObject(&gtd))

	m.EXPECT().DescribeGlobalTable(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&dynamodb.DescribeGlobalTableOutput{
			GlobalTableDescription: &gtd,
		},
		nil,
	)

	return services
}

func TestDynamodbGlobalTables(t *testing.T) {
	client.AwsMockTestHelper(t, GlobalTables(), buildDynamodbGlobalTablesMock, client.TestOptions{})
}
