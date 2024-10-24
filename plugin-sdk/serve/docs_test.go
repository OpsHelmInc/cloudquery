package serve

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/internal/memdb"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/plugin"
)

func TestPluginDocs(t *testing.T) {
	tmpDir := t.TempDir()
	p := plugin.NewPlugin(
		"testPlugin",
		"v1.0.0",
		memdb.NewMemDBClient)
	srv := Plugin(p)
	cmd := srv.newCmdPluginRoot()
	cmd.SetArgs([]string{"doc", tmpDir})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
}
