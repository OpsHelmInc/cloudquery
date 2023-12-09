package ec2

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
)

func buildRegionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	r := types.Region{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.OptInStatus = aws.String("opted-in")
	r.RegionName = aws.String("us-east-1")
	m.EXPECT().DescribeRegions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeRegionsOutput{
			Regions: []types.Region{r},
		}, nil)
	m.EXPECT().GetEbsDefaultKmsKeyId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsDefaultKmsKeyIdOutput{KmsKeyId: aws.String("some/key/id")}, nil)
	m.EXPECT().GetEbsEncryptionByDefault(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsEncryptionByDefaultOutput{EbsEncryptionByDefault: aws.Bool(true)}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestRegions(t *testing.T) {
	client.AwsMockTestHelper(t, Regions(), buildRegionsMock, client.TestOptions{})
}
