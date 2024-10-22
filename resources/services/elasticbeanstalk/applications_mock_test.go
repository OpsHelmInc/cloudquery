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

func buildElasticbeanstalkApplications(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	require.NoError(t, faker.FakeObject(&la))

	tag := elasticbeanstalkTypes.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	m.EXPECT().DescribeApplications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.DescribeApplicationsOutput{
			Applications: []elasticbeanstalkTypes.ApplicationDescription{la},
		}, nil)

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&elasticbeanstalk.ListTagsForResourceOutput{
			ResourceTags: []elasticbeanstalkTypes.Tag{tag},
			ResourceArn:  la.ApplicationArn,
		}, nil)

	return client.Services{
		Elasticbeanstalk: m,
	}
}

func TestElasticbeanstalkApplications(t *testing.T) {
	client.AwsMockTestHelper(t, Applications(), buildElasticbeanstalkApplications, client.TestOptions{})
}
