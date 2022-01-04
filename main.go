package main

import (
	_ "embed"
	"errors"
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/updater"
	"log"
	"os"
)

func main() {
	// self update service
	if len(os.Args) == 1 {
		updater.NewUpdaterService(packageJson).CheckForUpdates()
		// testing
		os.Exit(0);
	}

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

	// check if running in docker
	isInDocker := true
	if _, err := os.Stat("/.dockerenv"); errors.Is(err, os.ErrNotExist) {
		isInDocker = false
	}

	// ran via executable on desktop
	if len(os.Args) == 1 && !isInDocker {
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
	changelog string
	//go:embed package.json
	packageJson []byte
)

func loadEmbedded(app *boot.App) {
	// load embedded files into cache
	app.Cache().Set("changelog", changelog, -1)
	app.Cache().Set("packageJson", packageJson, -1)
}
