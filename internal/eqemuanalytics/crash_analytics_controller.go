package eqemuanalytics

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type CrashAnalyticsController struct {
	logger *logrus.Logger
	db     *database.DatabaseResolver
}

func NewCrashAnalyticsController(
	logger *logrus.Logger,
	db *database.DatabaseResolver,
) *CrashAnalyticsController {
	return &CrashAnalyticsController{
		logger: logger,
		db:     db,
	}
}

func (a *CrashAnalyticsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "server-crash-report", a.serverCrashReport, nil),
	}
}

func (a *CrashAnalyticsController) serverCrashReport(c echo.Context) error {
	r := new(models.CrashReport)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if r.Cpus == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.ServerName) == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if r.Uptime == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.ServerShortName) == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.CrashReport) == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.CrashReport) > 100000 || len(r.CrashReport) < 500 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if !strings.Contains(r.CrashReport, "stack trace for") &&
		!strings.Contains(r.CrashReport, "Windows") {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if strings.Count(r.ServerVersion, ".") != 2 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}

	a.db.GetSpireDb().Create(r)

	return c.JSON(http.StatusOK, "Invalid request")
}
