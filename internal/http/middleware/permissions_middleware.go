package middleware

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"net/http"
	"strings"
)

type PermissionsMiddleware struct {
	db          *database.Resolver
	cache       *gocache.Cache
	permissions *permissions.Service
	logger      *logger.AppLogger
}

func NewPermissionsMiddleware(
	db *database.Resolver,
	logger *logger.AppLogger,
	cache *gocache.Cache,
	permissions *permissions.Service,
) *PermissionsMiddleware {
	return &PermissionsMiddleware{
		db:          db,
		logger:      logger,
		cache:       cache,
		permissions: permissions,
	}
}

func (r PermissionsMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := request.GetUser(c)

			if user.ID > 0 {

				// key for the users database connection identifier
				connectionIdKey := fmt.Sprintf("active-connection-%v", user.ID)

				// find cached connection attempt #1
				// relies on database resolver logic to have ran so we trigger it once if not found
				cachedConn, found := r.cache.Get(connectionIdKey)
				if !found {
					r.db.Get(&models.DbStr{}, c) // trigger logic to hydrate cache
					cachedConn, found = r.cache.Get(connectionIdKey)
				}

				// found cached connection
				if found {
					connectionId := cachedConn.(uint)
					if connectionId > 0 {
						pass := r.permissions.CanAccessResource(c, user, connectionId)

						// get resource name
						params := strings.Split(c.Request().URL.Path, "/")
						resource := ""
						if len(params) > 0 {
							resource = strings.TrimSpace(params[3])
						}

						r.logger.DebugVvv().
							Any("user", user.ID).
							Any("connectionId", connectionId).
							Any("pass", pass).
							Any("resource", resource).
							Any("url", c.Request().URL.Path).
							Msg("permissions middleware")

						// did not pass ACL
						if !pass {
							return c.JSON(
								http.StatusForbidden,
								echo.Map{"error": fmt.Sprintf("You do not have permission to %v this resource [%v]", r.getAccessDescription(c), c.Request().URL.Path)},
							)
						}
					}
				}
			}

			return next(c)
		}
	}
}

func (r PermissionsMiddleware) getAccessDescription(c echo.Context) string {
	if r.permissions.IsWriteRequest(c) {
		return "write to"
	}

	return "read"
}
