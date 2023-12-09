package servicecatalog

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/golang/mock/gomock"
)

func buildProvisionedProducts(t *testing.T, ctrl *gomock.Controller) client.Services {
	mk := mocks.NewMockServicecatalogClient(ctrl)

	o := servicecatalog.SearchProvisionedProductsOutput{}
	if err := faker.FakeObject(&o); err != nil {
		t.Fatal(err)
	}
	o.NextPageToken = nil

	mk.EXPECT().SearchProvisionedProducts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&o,
		nil,
	)

	return client.Services{
		Servicecatalog: mk,
	}
}

func TestProvisionedProducts(t *testing.T) {
	client.AwsMockTestHelper(t, ProvisionedProducts(), buildProvisionedProducts, client.TestOptions{})
}
