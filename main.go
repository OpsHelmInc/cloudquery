package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/serve"
	internalPlugin "github.com/OpsHelmInc/cloudquery/v2/resources/plugin"
)

func main() {
	if err := serve.Plugin(internalPlugin.AWS(&aws.Config{})).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
