package plugin

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/plugins"
	"github.com/aws/aws-sdk-go-v2/aws"
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
