package eqemuserver

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DistributedZsController struct {
	db             *database.Resolver
	eqemuserverapi *Client
	pathmgmt       *pathmgmt.PathManagement
	serverconfig   *eqemuserverconfig.Config
	updater        *Updater
	launcher       *Launcher
}

func NewDistributedZsController(
	db *database.Resolver,
	api *Client,
	serverconfig *eqemuserverconfig.Config,
	pathmgmt *pathmgmt.PathManagement,
	launcher *Launcher,
) *DistributedZsController {
	return &DistributedZsController{
		db:             db,
		eqemuserverapi: api,
		serverconfig:   serverconfig,
		pathmgmt:       pathmgmt,
		launcher:       launcher,
	}
}

func (a *DistributedZsController) DzsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}

func (a *DistributedZsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "eqemuserver/dzs/test", a.test, []echo.MiddlewareFunc{a.DzsMiddleware()}),
	}
}

func (a *DistributedZsController) test(c echo.Context) error {

	return c.JSON(http.StatusOK, "hello")
}
