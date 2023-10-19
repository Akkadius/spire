package middleware

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type ReadOnlyMiddleware struct {
	db     *database.Resolver
	logger *logrus.Logger
}

func NewReadOnlyMiddleware(db *database.Resolver, logger *logrus.Logger) *ReadOnlyMiddleware {
	return &ReadOnlyMiddleware{
		db:     db,
		logger: logger,
	}
}

func (r ReadOnlyMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// continue request if this mode is not even enabled
			if !env.IsHostedReadOnlyModeEnabled() {
				return next(c)
			}

			// if it is a get request, we should be able to read it
			ignoredPostRoutes := []string{
				"/api/v1/connection",
				"/bulk",
			}

			if strings.Contains(c.Request().URL.Path, "/api/v1/connection") {
				return next(c)
			}

			// allow get calls
			// allow post calls that are ignored
			if c.Request().Method == "GET" || (c.Request().Method == "POST" && !contains(ignoredPostRoutes, c.Request().URL.Path)) {
				return next(c)
			}

			// anything else we assume is something attempting to write
			user := request.GetUser(c)
			if user.ID == 0 {
				return c.JSON(
					http.StatusForbidden,
					echo.Map{"error": "Not logged in"},
				)
			}

			// if we have a user, lets check to see if our user database is equal to the default
			if user.ID > 0 {
				userDb, err := r.db.ResolveUserEqemuConnection(&models.Zone{}, user).DB()
				if err != nil {
					return c.JSON(
						http.StatusForbidden,
						echo.Map{"error": fmt.Sprintf("[read_only_middleware] Can't get user database %v", err.Error())},
					)
				}
				eqemuDb, err := r.db.GetEqemuDb().DB()
				if err != nil {
					return c.JSON(
						http.StatusForbidden,
						echo.Map{"error": fmt.Sprintf("[read_only_middleware] Can't get eqemu database %v", err.Error())},
					)
				}

				if userDb == eqemuDb {
					r.logger.Debugf("Database instance is equal, read only mode")

					return c.JSON(
						http.StatusForbidden,
						echo.Map{"error": "Database is in read only mode"},
					)
				}
			}

			return next(c)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
