package main

import (
	"context"
	"log"

	internalPlugin "github.com/OpsHelmInc/cloudquery/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	if err := serve.Plugin(internalPlugin.AWS()).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}
