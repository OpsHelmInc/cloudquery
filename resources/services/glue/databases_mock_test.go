package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildDatabasesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	db := glue.GetDatabasesOutput{}
	require.NoError(t, faker.FakeObject(&db))
	db.NextToken = nil
	m.EXPECT().GetDatabases(gomock.Any(), gomock.Any()).Return(&db, nil)

	tb := glue.GetTablesOutput{}
	require.NoError(t, faker.FakeObject(&tb))
	tb.NextToken = nil
	m.EXPECT().GetTables(gomock.Any(), gomock.Any()).Return(&tb, nil)

	i := glue.GetPartitionIndexesOutput{}
	require.NoError(t, faker.FakeObject(&i))
	i.NextToken = nil
	m.EXPECT().GetPartitionIndexes(gomock.Any(), gomock.Any()).Return(&i, nil)

	// get database tags
	tags := glue.GetTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().GetTags(gomock.Any(), gomock.Any()).Return(&tags, nil)

	// get table tags
	tableTags := glue.GetTagsOutput{}
	require.NoError(t, faker.FakeObject(&tableTags))
	m.EXPECT().GetTags(gomock.Any(), gomock.Any()).Return(&tableTags, nil)

	return client.Services{
		Glue: m,
	}
}

func TestDatabases(t *testing.T) {
	client.AwsMockTestHelper(t, Databases(), buildDatabasesMock, client.TestOptions{})
}
