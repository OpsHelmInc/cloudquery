package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildCognitoIdentityPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoidentityClient(ctrl)

	var desc types.IdentityPoolShortDescription
	require.NoError(t, faker.FakeObject(&desc))

	m.EXPECT().ListIdentityPools(
		gomock.Any(),
		&cognitoidentity.ListIdentityPoolsInput{MaxResults: aws.Int32(60)},
		gomock.Any(),
	).Return(
		&cognitoidentity.ListIdentityPoolsOutput{IdentityPools: []types.IdentityPoolShortDescription{desc}},
		nil,
	)

	var ipo cognitoidentity.DescribeIdentityPoolOutput
	require.NoError(t, faker.FakeObject(&ipo))

	ipo.IdentityPoolId = desc.IdentityPoolId
	ipo.IdentityPoolId = desc.IdentityPoolName
	m.EXPECT().DescribeIdentityPool(
		gomock.Any(),
		&cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: desc.IdentityPoolId},
		gomock.Any(),
	).Return(&ipo, nil)

	return client.Services{Cognitoidentity: m}
}

func TestCognitoIdentityPools(t *testing.T) {
	client.AwsMockTestHelper(t, IdentityPools(), buildCognitoIdentityPools, client.TestOptions{})
}
