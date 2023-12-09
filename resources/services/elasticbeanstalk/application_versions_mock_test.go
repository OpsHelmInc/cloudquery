package elasticbeanstalk

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/golang/mock/gomock"
)

func buildElasticbeanstalkApplicationVersions(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationVersionDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)
	}

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
