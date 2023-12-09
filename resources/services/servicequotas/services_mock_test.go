package servicequotas

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/faker"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/golang/mock/gomock"
)

func buildServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockServicequotasClient(ctrl)

	services := servicequotas.ListServicesOutput{}
	err := faker.FakeObject(&services)
	if err != nil {
		t.Fatal(err)
	}
	services.NextToken = nil
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).Return(&services, nil)

	quotas := servicequotas.ListServiceQuotasOutput{}
	err = faker.FakeObject(&quotas)
	if err != nil {
		t.Fatal(err)
	}

	quotas.NextToken = nil
	m.EXPECT().ListServiceQuotas(gomock.Any(), gomock.Any(), gomock.Any()).Return(&quotas, nil).AnyTimes()

	return client.Services{
		Servicequotas: m,
	}
}

func TestQuotas(t *testing.T) {
	client.AwsMockTestHelper(t, Services(), buildServices, client.TestOptions{})
}
