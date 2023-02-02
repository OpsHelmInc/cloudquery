package plugin

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (
	Version = "Development"
)

func AWS(config aws.Config) *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"aws",
		Version,
		tables(),
		client.CustomConfig(config),
	)
}
