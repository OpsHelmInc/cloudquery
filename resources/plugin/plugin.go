package plugin

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (
	Version = "Development"
)

func AWS() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
	)
}
