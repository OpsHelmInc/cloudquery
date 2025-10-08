package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/rs/zerolog/log"

	"github.com/OpsHelmInc/cloudquery/resources/plugin"
)

const sentryDSN = "https://sentryDSNEndpoint"

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal().Err(err)
	}
	serve.Source(plugin.AWS(cfg), serve.WithSourceSentryDSN(sentryDSN))
}
