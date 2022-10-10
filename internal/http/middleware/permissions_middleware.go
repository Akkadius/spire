package middleware

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PermissionsMiddleware struct {
	db     *database.DatabaseResolver
	cache  *gocache.Cache
	logger *logrus.Logger
}

func NewPermissionsMiddleware(db *database.DatabaseResolver, logger *logrus.Logger, cache *gocache.Cache) *PermissionsMiddleware {
	return &PermissionsMiddleware{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

func (r PermissionsMiddleware) Handle() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := request.GetUser(c)

			if ctx.ID > 0 {
				r.logger.Printf("Permissions middleware user [%v]\n", ctx.ID)

				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User does not have permission to access this resource"})
			}

			return next(c)
		}
	}
}
