package occulus

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
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

func (a *Controller) occulusProxy(c echo.Context) error {
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
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	}

	return c.String(
		response.StatusCode,
		string(body),
	)
}
