package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCloudfrontDistributionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	ds := cloudfrontTypes.DistributionSummary{}
	require.NoError(t, faker.FakeObject(&ds))

	cloudfrontOutput := &cloudfront.ListDistributionsOutput{
		DistributionList: &cloudfrontTypes.DistributionList{
			Items: []cloudfrontTypes.DistributionSummary{ds},
		},
	}
	m.EXPECT().ListDistributions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)

	distribution := &cloudfront.GetDistributionOutput{}
	require.NoError(t, faker.FakeObject(&distribution))

	m.EXPECT().GetDistribution(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		distribution,
		nil,
	)

	tags := &cloudfront.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestCloudfrontDistributions(t *testing.T) {
	client.AwsMockTestHelper(t, Distributions(), buildCloudfrontDistributionsMock, client.TestOptions{})
}
