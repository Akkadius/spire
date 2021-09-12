package main

import (
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/console"
	"github.com/Akkadius/spire/env"
	"log"
	"os"
	"runtime"
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

	// ran via executable on desktop
	if len(os.Args) == 1 && runtime.GOOS == "windows" {
		app.Desktop().Boot()
	}

	// run cmd
	if err := console.Run(app.Commands()); err != nil {
		log.Fatal(err)
	}
}
