package main

import (
	_ "embed"
	"fmt"
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/updater"
	"log"
	"os"
	"time"
)

func main() {
	// self update service
	if len(os.Args) == 1 && len(os.Getenv("SKIP_UPDATE_CHECK")) == 0 {
		updater.NewService(packageJson).CheckForUpdates()
	}

	// default
	_ = os.Setenv("APP_ENV", "local")

	// installer logic
	// check if eqemu_config.json exists in current directory
	// if it doesn't exist, run the eqemu server installer
	if _, err := os.Stat("eqemu_config.json"); os.IsNotExist(err) {
		eqemuserver.NewInstaller().Install()
		return
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

	// only hang if we're not running a command
	if len(os.Args) == 1 {
		fmt.Print("Automatically shutting down in 10 seconds...")
		time.Sleep(10 * time.Second)
		os.Exit(1)
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
