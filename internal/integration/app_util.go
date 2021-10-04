package integration

import (
	"github.com/Akkadius/spire/boot"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/models"
	"log"
)

var appBooted = false
var app *boot.App

func bootApp() *boot.App {
	if appBooted {
		return app
	}

	// load env
	if err := env.LoadEnvFileIfExists(); err != nil {
		log.Fatal(err)
	}

	// boot app
	booted, err := boot.InitializeApplication()
	if err != nil {
		log.Fatal(err)
	}

	appBooted = true
	app = &booted

	return &booted
}

var spireTables = []models.Modelable{
	&models.User{},
	&models.UserServerDatabaseConnection{},
	&models.ServerDatabaseConnection{},
}
