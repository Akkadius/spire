package middleware

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type LocalUserAuthMiddleware struct {
	db        *database.Resolver
	cache     *gocache.Cache
	logger    *logrus.Logger
	debug     int
	spireInit *spire.Init
	settings  *spire.Settings
}

func NewLocalUserAuthMiddleware(
	db *database.Resolver,
	logger *logrus.Logger,
	cache *gocache.Cache,
	settings *spire.Settings,
	spireInit *spire.Init,
) *LocalUserAuthMiddleware {
	return &LocalUserAuthMiddleware{
		db:        db,
		logger:    logger,
		cache:     cache,
		settings:  settings,
		spireInit: spireInit,
		debug:     env.GetInt("DEBUG", "0"),
	}
}

func (r LocalUserAuthMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := request.GetUser(c)

			if user.ID == 0 &&
				r.spireInit.IsInitialized() &&
				!strings.Contains(c.Request().RequestURI, "admin/occulus/download/") &&
				r.settings.IsSettingEnabled(spire.SettingAuthEnabled) {
				return c.JSON(
					http.StatusUnauthorized,
					echo.Map{"error": fmt.Sprintf("You are not logged in")},
				)
			}

			return next(c)
		}
	}
}
