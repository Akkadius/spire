package permissions

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	db          *database.Resolver
	permissions *Service
}

func NewController(
	db *database.Resolver,
	permissions *Service,
) *Controller {
	return &Controller{
		db:          db,
		permissions: permissions,
	}
}

func (p *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "permissions/resources", p.getPermissionResources, nil),
	}
}

func (p *Controller) getPermissionResources(c echo.Context) error {
	return c.JSON(http.StatusOK, p.permissions.GetResources(c.Echo().Routes()))
}
