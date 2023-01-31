package main

import (
	"github.com/OpsHelmInc/cloudquery/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "https://6c6b72bc946844cb8471f49eba485cde@o1396617.ingest.sentry.io/6747636"

func main() {
	serve.Source(plugin.AWS(), serve.WithSourceSentryDSN(sentryDSN))
}
