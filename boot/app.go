package boot

import (
	"context"
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/jinzhu/gorm"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// App is the root App resource
type App struct {
	context       context.Context
	mysql         *gorm.DB
	dbConnections *database.Connections
	logger        *logrus.Logger
	cache         *gocache.Cache
	commands      []*cobra.Command
	db            *database.DatabaseResolver
	router        *routes.Router
	desktop       *desktop.WebBoot
}

func (a App) Desktop() *desktop.WebBoot {
	return a.desktop
}

func (a App) DbConnections() *database.Connections {
	return a.dbConnections
}

func (a App) Commands() []*cobra.Command {
	return a.commands
}

// Create new App
func NewApplication(
	mysql *gorm.DB,
	logger *logrus.Logger,
	cache *gocache.Cache,
	commands []*cobra.Command,
	db *database.DatabaseResolver,
	dbConnections *database.Connections,
	router *routes.Router,
	desktop *desktop.WebBoot,
) App {
	return App{
		context:       context.Background(),
		mysql:         mysql,
		logger:        logger,
		cache:         cache,
		commands:      commands,
		db:            db,
		dbConnections: dbConnections,
		router:        router,
		desktop:       desktop,
	}
}
