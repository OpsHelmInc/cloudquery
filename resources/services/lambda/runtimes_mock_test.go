package lambda

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/golang/mock/gomock"
)

func buildLambdaRuntimesMock(*testing.T, *gomock.Controller) client.Services {
	return client.Services{}
}

func TestLambdaRuntimes(t *testing.T) {
	client.AwsMockTestHelper(t, Runtimes(), buildLambdaRuntimesMock, client.TestOptions{})
}
