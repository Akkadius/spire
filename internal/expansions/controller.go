package expansions

import (
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "expansions", a.get, nil),
	}
}

func (a *Controller) get(c echo.Context) error {
	return c.JSON(http.StatusOK, expansions)
}
