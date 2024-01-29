package main

import (
	_ "embed"
	"fmt"
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/updater"
	"github.com/henvic/httpretty"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// uncomment for profiling
	//f, err := os.Create("main.prof")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	registerHttpLogging()

	// self update service
	if len(os.Args) == 1 {
		if updater.NewService(packageJson).CheckForUpdates(true) {
			os.Exit(0)
		}
	}

	// default
	_ = os.Setenv("APP_ENV", "local")

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

func registerHttpLogging() {
	if len(os.Getenv("HTTP_DEBUG")) > 0 {
		logger := &httpretty.Logger{
			Time:           true,
			TLS:            true,
			RequestHeader:  true,
			RequestBody:    true,
			ResponseHeader: true,
			ResponseBody:   true,
			Colors:         true,
			Formatters:     []httpretty.Formatter{&httpretty.JSONFormatter{}},
		}

		http.DefaultClient.Transport = logger.RoundTripper(http.DefaultClient.Transport)
	}
}
