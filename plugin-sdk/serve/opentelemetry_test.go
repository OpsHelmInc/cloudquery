package serve

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/internal/memdb"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
)

func TestNewResource(t *testing.T) {
	p := plugin.NewPlugin(
		"testPluginV3",
		"v1.0.0",
		memdb.NewMemDBClient)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("failed to instantiate resource: %v", r)
		}
	}()
	newResource(p)
}
