package occulus

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Controller struct {
	db      *database.DatabaseResolver
	logger  *logrus.Logger
	occulus *Proxy
}

func NewController(
	logger *logrus.Logger,
	db *database.DatabaseResolver,
	occulus *Proxy,
) *Controller {
	return &Controller{
		logger:  logger,
		occulus: occulus,
		db:      db,
	}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "admin/test", a.test, nil),
		routes.RegisterRoute("ANY", "admin/occulus/*", a.occulusProxy, nil),
	}
}

type AdminEventRequest struct {
	EventName  string   `json:"event_name"`
	EventValue string   `json:"event_value"`
	Tags       []string `json:"tags"`
	Values     []string `json:"values"`
}

// searches quest examples
func (a *Controller) test(c echo.Context) error {

	// result
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": "Ok",
		},
	)
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func contains(val string, slice []string) bool {
	for _, item := range slice {
		if strings.Contains(val, item) {
			return true
		}
	}
	return false
}

func (a *Controller) occulusProxy(c echo.Context) error {
	// occulus is almost entirely removed as a dependency
	// only allow routes that we need to at this point to reduce
	// security concerns for bad actors
	allowedRoutes := []string{
		"git/quests/branch",
		"server/launcher/config",
		"server/restart",
		"server/start",
		"server/stop",
		"server/websocket-authorization",
	}
	if !contains(c.Request().URL.Path, allowedRoutes) {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "This Occulus route is not allowed"})
	}

	response, body, err := a.occulus.ProxyRequest(c)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": err.Error(),
			},
		)
	}

	if response != nil {
		// delete the existing headers
		for header := range c.Response().Header() {
			c.Response().Header().Del(header)
		}
		for key, values := range response.Header {
			for _, value := range values {
				c.Response().Header().Add(key, value)
			}
		}
		c.Response().Header().Del("Access-Control-Allow-Origin")
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	}

	return c.String(
		response.StatusCode,
		string(body),
	)
}
