package controllers

import (
	"github.com/Akkadius/spire/http/request"
	"github.com/Akkadius/spire/http/routes"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AnalyticsController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	influx *influx.Client
}

func NewAnalyticsController(
	logger *logrus.Logger,
	influx *influx.Client,
	db *database.DatabaseResolver,
) *AnalyticsController {
	return &AnalyticsController{logger: logger, influx: influx, db: db}
}

func (a *AnalyticsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "analytics/event", a.event, nil),
		routes.RegisterRoute(http.MethodPost, "analytics/count", a.count, nil),
	}
}

type AnalyticsEventRequest struct {
	EventName  string   `json:"event_name"`
	EventValue string   `json:"event_value"`
	Tags       []string `json:"tags"`
	Values     []string `json:"values"`
}

// searches quest examples
func (a *AnalyticsController) event(c echo.Context) error {
	if a.db.GetSpireDb() == nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": "Spire database unavailable",
			},
		)
	}

	// body - bind
	p := new(AnalyticsEventRequest)
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

type AnalyticsEventCountRequest struct {
	EventName string `json:"event_name"`
	EventKey  string `json:"event_key"`
}

// searches quest examples
func (a *AnalyticsController) count(c echo.Context) error {
	if a.db.GetSpireDb() == nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{
				"message": "Spire database unavailable",
			},
		)
	}

	r := new(AnalyticsEventCountRequest)
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
