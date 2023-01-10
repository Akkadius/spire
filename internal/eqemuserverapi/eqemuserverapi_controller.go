package eqemuserverapi

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	db             *database.DatabaseResolver
	logger         *logrus.Logger
	eqemuserverapi *Client
}

func NewController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	api *Client,
) *Controller {
	return &Controller{
		db:             db,
		logger:         logger,
		eqemuserverapi: api,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "eqemuserver/reload-types", a.getReloadTypes, nil),
		routes.RegisterRoute(http.MethodPost, "eqemuserver/reload/:type", a.reload, nil),
	}
}

func (a *Controller) getReloadTypes(c echo.Context) error {
	types, err := a.eqemuserverapi.GetReloadTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, types)
}

func (a *Controller) reload(c echo.Context) error {
	reloadType := c.Param("type")
	r, err := a.eqemuserverapi.Reload(reloadType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, r)
}
