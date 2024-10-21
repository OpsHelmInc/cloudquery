package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCrawlers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGlueClient(ctrl)

	var crawler glue.GetCrawlersOutput
	require.NoError(t, faker.FakeObject(&crawler))

	crawler.NextToken = nil
	m.EXPECT().GetCrawlers(gomock.Any(), gomock.Any(), gomock.Any()).Return(&crawler, nil)

	var tags glue.GetTagsOutput
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return client.Services{
		Glue: m,
	}
}

func TestCrawlers(t *testing.T) {
	client.AwsMockTestHelper(t, Crawlers(), buildCrawlers, client.TestOptions{})
}
