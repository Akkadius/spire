package eqemuanalytics

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
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

type CrashReport struct {
	CompileDate     string  `json:"compile_date"`
	CompileTime     string  `json:"compile_time"`
	Cpus            int     `json:"cpus"`
	CrashReport     string  `json:"crash_report"`
	OsMachine       string  `json:"os_machine"`
	OsRelease       string  `json:"os_release"`
	OsSysname       string  `json:"os_sysname"`
	OsVersion       string  `json:"os_version"`
	ProcessID       int     `json:"process_id"`
	RssMemory       float64 `json:"rss_memory"`
	ServerName      string  `json:"server_name"`
	ServerShortName string  `json:"server_short_name"`
	ServerVersion   string  `json:"server_version"`
	Uptime          int     `json:"uptime"`
}

func (a *CrashAnalyticsController) serverCrashReport(c echo.Context) error {
	r := new(CrashReport)
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

	return c.JSON(http.StatusOK, "Invalid request")
}
