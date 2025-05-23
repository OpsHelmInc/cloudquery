package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildDynamodbTablesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDynamodbClient(ctrl)
	services := client.Services{
		Dynamodb: m,
	}
	var tableName string
	require.NoError(t, faker.FakeObject(&tableName))

	listOutput := &dynamodb.ListTablesOutput{
		TableNames: []string{tableName},
	}
	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listOutput,
		nil,
	)

	descOutput := &dynamodb.DescribeTableOutput{
		Table: &types.TableDescription{
			TableName:          &tableName,
			GlobalTableVersion: aws.String("2019.11.21"),
		},
	}
	require.NoError(t, faker.FakeObject(descOutput.Table))

	m.EXPECT().DescribeTable(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		descOutput,
		nil,
	)

	repOutput := &dynamodb.DescribeTableReplicaAutoScalingOutput{
		TableAutoScalingDescription: &types.TableAutoScalingDescription{
			TableName:   &tableName,
			TableStatus: types.TableStatusActive,
		},
	}
	require.NoError(t, faker.FakeObject(&repOutput.TableAutoScalingDescription.Replicas))

	m.EXPECT().DescribeTableReplicaAutoScaling(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		repOutput,
		nil,
	)

	cbOutput := &dynamodb.DescribeContinuousBackupsOutput{
		ContinuousBackupsDescription: &types.ContinuousBackupsDescription{},
	}
	require.NoError(t, faker.FakeObject(&cbOutput.ContinuousBackupsDescription))

	m.EXPECT().DescribeContinuousBackups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cbOutput,
		nil,
	)

	tags := &dynamodb.ListTagsOfResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	tags.NextToken = nil
	m.EXPECT().ListTagsOfResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestDynamodbTables(t *testing.T) {
	client.AwsMockTestHelper(t, Tables(), buildDynamodbTablesMock, client.TestOptions{})
}
