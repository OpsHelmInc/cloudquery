package servicecatalog

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry"
	"github.com/golang/mock/gomock"
)

func buildPortfolios(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)
	ma := mocks.NewMockServicecatalogappregistryClient(ctrl)

	o := servicecatalog.ListPortfoliosOutput{}
	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)
	}
	o.NextPageToken = nil

	mk.EXPECT().ListPortfolios(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	to := servicecatalogappregistry.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&to); err != nil {
		t.Fatal(err)
	}

	ma.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&to,
		nil,
	)

	return client.Services{
		Servicecatalog:            mk,
		Servicecatalogappregistry: ma,
	}
}

func TestPortfolios(t *testing.T) {
	client.AwsMockTestHelper(t, Portfolios(), buildPortfolios, client.TestOptions{})
}
