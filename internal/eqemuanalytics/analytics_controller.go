package eqemuanalytics

import (
	"github.com/Akkadius/spire/internal/database"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AnalyticsController struct {
	logger   *logrus.Logger
	db       *database.DatabaseResolver
	releases *Releases
}

func NewAnalyticsController(
	logger *logrus.Logger,
	db *database.DatabaseResolver,
	releases *Releases,
) *AnalyticsController {
	return &AnalyticsController{
		logger:   logger,
		db:       db,
		releases: releases,
	}
}

func v1AnalyticsRateLimit() echo.MiddlewareFunc {
	return appmiddleware.RateLimiterWithConfig(
		appmiddleware.RateLimiterConfig{
			LimitConfig: appmiddleware.LimiterConfig{
				Max:      15,
				Duration: time.Minute * 1,
				Strategy: "ip",
				Key:      "",
			},
			Prefix:                       "ANALYTICS-LIMIT",
			Client:                       nil,
			SkipRateLimiterInternalError: false,
			OnRateLimit:                  nil,
		},
	)
}

func (a *AnalyticsController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "analytics/server-crash-report", a.serverCrashReport, []echo.MiddlewareFunc{v1AnalyticsRateLimit()}),
		routes.RegisterRoute(http.MethodGet, "analytics/server-crash-reports", a.listServerCrashReports, nil),
		routes.RegisterRoute(http.MethodGet, "analytics/server-crash-report/counts", a.getServerCrashReportCounts, nil),
		routes.RegisterRoute(http.MethodGet, "analytics/releases", a.getReleases, nil),
	}
}

func (a *AnalyticsController) serverCrashReport(c echo.Context) error {
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
	if len(r.ServerShortName) == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.CrashReport) == 0 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if len(r.CrashReport) > 100000 || len(r.CrashReport) < 500 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if !strings.Contains(r.CrashReport, "print_trace") &&
		!strings.Contains(r.CrashReport, "Windows") {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}
	if strings.Count(r.ServerVersion, ".") != 2 {
		return c.JSON(http.StatusInternalServerError, "Invalid request")
	}

	a.db.GetSpireDb().Create(r)

	return c.JSON(http.StatusOK, "Invalid request")
}

func (a *AnalyticsController) listServerCrashReports(c echo.Context) error {
	var entries []models.CrashReport
	q := a.db.GetSpireDb()

	// limit
	paramLimit := c.QueryParam("limit")
	limit := 1000
	if len(paramLimit) > 0 {
		l, err := strconv.Atoi(paramLimit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		limit = l
	}
	q = q.Limit(limit)

	// version filters
	w := c.QueryParam("version")
	if len(w) > 0 {
		q = q.Where("server_version", w)
	}

	// paging
	queryParamOffset := c.QueryParam("page")
	if len(queryParamOffset) > 0 {
		queryOffset, err := strconv.Atoi(queryParamOffset)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		q = q.Offset(queryOffset * limit)
	}

	q.Find(&entries)

	return c.JSON(http.StatusOK, entries)
}

type CrashReportCounts struct {
	ServerVersion string `json:"server_version"`
	CompileDate   string `json:"compile_date"`
	CrashCount    int    `json:"crash_count"`
}

func (a *AnalyticsController) getServerCrashReportCounts(c echo.Context) error {
	db, err := a.db.GetSpireDb().DB()
	if err != nil {
		return err
	}

	rows, err := db.Query("select server_version, compile_date, count(*) as crash_count from spire_crash_reports group by server_version order by server_version")
	if err != nil {
		return err
	}

	counts := []CrashReportCounts{}
	for rows.Next() {
		var r CrashReportCounts
		err = rows.Scan(&r.ServerVersion, &r.CompileDate, &r.CrashCount)
		if err != nil {
			return err
		}
		counts = append(counts, r)
	}

	return c.JSON(http.StatusOK, counts)
}

func (a *AnalyticsController) getReleases(c echo.Context) error {
	r, err := a.releases.getReleases()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{"data": r},
	)
}
