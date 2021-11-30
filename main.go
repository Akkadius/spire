package main

import (
	_ "embed"
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/console"
	"github.com/Akkadius/spire/internal/env"
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

	// load embedded resources
	loadEmbedded(&app)

	// ran via executable on desktop
	if len(os.Args) == 1 && runtime.GOOS == "windows" {
		_ = os.Setenv("APP_ENV", "desktop")
		app.Desktop().Boot()
	}

	// run cmd
	if err := console.Run(app.Commands()); err != nil {
		log.Fatal(err)
	}
}

var (
	//go:embed CHANGELOG.md
	changelog   string
	//go:embed package.json
	packageJson []byte
)

func loadEmbedded(app *boot.App) {
	// load embedded files into cache
	app.Cache().Set("changelog", changelog, -1)
	app.Cache().Set("packageJson", packageJson, -1)
}
