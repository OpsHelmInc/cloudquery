package servicecatalog

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/faker"
)

func buildPortfolios(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)

	o := servicecatalog.ListPortfoliosOutput{}
	require.NoError(t, faker.FakeObject(&o))

	o.NextPageToken = nil

	mk.EXPECT().ListPortfolios(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	po := servicecatalog.DescribePortfolioOutput{}
	require.NoError(t, faker.FakeObject(&po))

	mk.EXPECT().DescribePortfolio(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&po,
		nil,
	)

	return client.Services{
		Servicecatalog: mk,
	}
}

func TestPortfolios(t *testing.T) {
	client.AwsMockTestHelper(t, Portfolios(), buildPortfolios, client.TestOptions{})
}
