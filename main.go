package main

import (
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/console"
	"github.com/Akkadius/spire/env"
	"log"
)

func main() {
	// load env
	if err := env.LoadEnvFileIfExists(); err != nil {
		log.Fatal(err)
	}

	// boot app
	app, err := boot.InitializeApplication()
	if err != nil {
		log.Fatal(err)
	}

	// run cmd
	if err := console.Run(app.Commands()); err != nil {
		log.Fatal(err)
	}
}
