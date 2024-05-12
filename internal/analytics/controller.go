package analytics

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Controller struct {
	db     *database.Resolver
	influx *influx.Client
}

func NewController(
	influx *influx.Client,
	db *database.Resolver,
) *Controller {
	return &Controller{influx: influx, db: db}
}

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "analytics/event", a.event, nil),
		routes.RegisterRoute(http.MethodPost, "analytics/count", a.count, nil),
	}
}

type EventRequest struct {
	EventName  string   `json:"event_name"`
	EventValue string   `json:"event_value"`
	Tags       []string `json:"tags"`
	Values     []string `json:"values"`
}

// searches quest examples
func (a *Controller) event(c echo.Context) error {
	if a.db.GetSpireDb() == nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": "Spire database unavailable",
			},
		)
	}

	// body - bind
	p := new(EventRequest)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := request.GetUser(c)
	event := models.AnalyticEvent{
		EventName:  p.EventName,
		EventValue: p.EventValue,
		RequestUri: c.Request().RequestURI,
		IpAddress:  c.RealIP(),
		UserID:     user.ID,
		CreatedAt:  time.Time{},
	}
	a.db.GetSpireDb().Create(&event)

	// result
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": "Ok",
		},
	)
}

type EventCountRequest struct {
	EventName string `json:"event_name"`
	EventKey  string `json:"event_key"`
}

// searches quest examples
func (a *Controller) count(c echo.Context) error {
	if a.db.GetSpireDb() == nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": "Spire database unavailable",
			},
		)
	}

	r := new(EventCountRequest)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	event := models.AnalyticEventCount{
		EventName: r.EventName,
		EventKey:  r.EventKey,
	}

	a.db.GetSpireDb().Where(event).First(&event)
	if event.ID > 0 {
		a.db.GetSpireDb().Model(&event).Where(event).Update("count", event.Count+1)
	}

	if event.ID == 0 {
		a.db.GetSpireDb().Model(&event).Create(&event)
	}

	// result
	return c.JSON(
		http.StatusOK,
		echo.Map{
			"data": "Ok",
		},
	)
}
