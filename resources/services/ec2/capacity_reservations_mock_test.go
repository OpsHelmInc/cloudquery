package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2CapacityReservations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.CapacityReservation{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeCapacityReservations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeCapacityReservationsOutput{
			CapacityReservations: []types.CapacityReservation{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2CapacityReservations(t *testing.T) {
	client.AwsMockTestHelper(t, CapacityReservations(), buildEc2CapacityReservations, client.TestOptions{})
}
