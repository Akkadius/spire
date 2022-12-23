package main

import (
	"bufio"
	_ "embed"
	"fmt"
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
	}

	// load env
	if err := env.LoadEnvFileIfExists(); err != nil {
		Fatal(err)
	}

	// boot app
	app, err := boot.InitializeApplication()
	if err != nil {
		Fatal(err)
	}

	// load embedded resources
	loadEmbedded(&app)

	// ran via executable on desktop
	if len(os.Args) == 1 {
		_ = os.Setenv("APP_ENV", "desktop")
		app.Desktop().Boot()
	}

	// run cmd
	if err := console.Run(app.Commands()); err != nil {
		Fatal(err)
	}
}

func Fatal(err error) {
	log.Println(err)
	fmt.Print("Press any to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
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
