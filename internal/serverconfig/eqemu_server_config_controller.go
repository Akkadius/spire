package serverconfig

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	logger       *logrus.Logger
	serverconfig *EQEmuServerConfig
}

func NewController(
	logger *logrus.Logger,
	serverconfig *EQEmuServerConfig,
) *Controller {
	return &Controller{
		logger:       logger,
		serverconfig: serverconfig,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "admin/serverconfig", a.get, nil),
		routes.RegisterRoute(http.MethodPost, "admin/serverconfig", a.save, nil),
	}
}

func (a *Controller) get(c echo.Context) error {
	return c.JSON(http.StatusOK, a.serverconfig.Get())
}

func (a *Controller) save(c echo.Context) error {
	var config EQEmuConfigJson
	err := c.Bind(&config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to bind config")
	}

	if len(config.Server.World.Longname) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Configuration long name is empty!"})
	}

	if len(config.Server.World.Shortname) == 0 {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Configuration short name is empty!"})
	}

	err = a.serverconfig.Save(config)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, config)
}
