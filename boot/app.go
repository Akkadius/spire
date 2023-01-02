package boot

import (
	"context"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/spire"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
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
	onboarding    *spire.Init
}

func (a App) Cache() *gocache.Cache {
	return a.cache
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
	onboarding *spire.Init,
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
		onboarding:    onboarding,
	}
}
