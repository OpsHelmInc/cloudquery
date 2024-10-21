package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildEc2Instances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.Reservation{}
	require.NoError(t, faker.FakeObject(&l))

	l.Instances[0].StateTransitionReason = aws.String("User initiated (2021-11-26 11:33:00 GMT)")
	creationDate := "1994-11-05T08:15:30-05:00"
	l.Instances[0].ElasticGpuAssociations[0].ElasticGpuAssociationTime = &creationDate
	nextToken := "test"
	// this test ensures that pagination works by returning a token once, but not the second time
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []types.Reservation{},
			NextToken:    &nextToken,
		}, nil)
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []types.Reservation{l},
			NextToken:    nil,
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2Instances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildEc2Instances, client.TestOptions{})
}
