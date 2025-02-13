package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildElasticbeanstalkApplicationVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationVersionDescription{}
	require.NoError(t, faker.FakeObject(&la))

	m.EXPECT().DescribeApplicationVersions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeApplicationVersionsOutput{
			ApplicationVersions: []elasticbeanstalkTypes.ApplicationVersionDescription{la},
		}, nil)

	return client.Services{
		Elasticbeanstalk: m,
	}
}

func TestElasticbeanstalkApplicationVersions(t *testing.T) {
	client.AwsMockTestHelper(t, ApplicationVersions(), buildElasticbeanstalkApplicationVersions, client.TestOptions{})
}
