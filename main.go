package main

import (
	"context"
	"log"

	internalPlugin "github.com/OpsHelmInc/cloudquery/resources/plugin"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	if err := serve.Plugin(internalPlugin.AWS(&aws.Config{})).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
