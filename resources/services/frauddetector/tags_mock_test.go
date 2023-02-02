package frauddetector

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func addTagsCall(t *testing.T, client *mocks.MockFrauddetectorClient) {
	var data []types.Tag
	err := faker.FakeObject(&data)
	if err != nil {
		t.Fatal(err)
	}

	client.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&frauddetector.ListTagsForResourceOutput{Tags: data}, nil,
	)
}
