package eqemuanalytics

import (
	"fmt"
	"github.com/Akkadius/spire/internal/crashreporting"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/discord"
	appmiddleware "github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Controller struct {
	logger   *logrus.Logger
	db       *database.Resolver
	releases *Releases
}

func NewController(
	logger *logrus.Logger,
	db *database.Resolver,
	releases *Releases,
) *Controller {
	return &Controller{
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

func (a *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodPost, "analytics/server-crash-report", a.serverCrashReport, []echo.MiddlewareFunc{v1AnalyticsRateLimit()}),
		routes.RegisterRoute(http.MethodGet, "analytics/server-crash-reports", a.listServerCrashReports, nil),
		routes.RegisterRoute(http.MethodGet, "analytics/server-crash-report/counts", a.getServerCrashReportCounts, nil),
		routes.RegisterRoute(http.MethodGet, "analytics/releases", a.getReleases, nil),
	}
}

func (a *Controller) serverCrashReport(c echo.Context) error {
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

	r.Fingerprint = crashreporting.FingerPrint(r.CrashReport)

	// check database if fingerprint exists
	count := int64(-1)
	isReleaseVersion := !strings.Contains(r.ServerVersion, "-dev")

	if isReleaseVersion {
		a.db.GetSpireDb().
			Model(&models.CrashReport{}).
			Where("fingerprint = ?", r.Fingerprint).
			Count(&count)
	}

	a.db.GetSpireDb().Create(r)

	// if count is 0, then this is a new crash
	if count == 0 && isReleaseVersion {
		// send discord webhook
		go func() {

			// format web link
			link := fmt.Sprintf(
				"%s/dev/release/%s?id=%d",
				os.Getenv("VUE_APP_FRONTEND_BASE_URL"),
				r.ServerVersion,
				r.ID,
			)

			err := discord.SendDiscordWebhook(
				os.Getenv("DISCORD_CRASH_REPORT_WEBHOOK_URL"),
				fmt.Sprintf(
					"Version **%v** New crash fingerprint **%v** created by server **%v** can be viewed at %v",
					r.ServerVersion,
					r.Fingerprint,
					r.ServerName,
					link,
				),
			)
			if err != nil {
				a.logger.Error(err)
			}
		}()
	}

	return c.JSON(http.StatusOK, "Invalid request")
}

func (a *Controller) listServerCrashReports(c echo.Context) error {
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
	q = q.Preload("ResolvedUser")

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

type CrashReportResponse struct {
	CrashReportCounts []CrashReportCounts       `json:"crash_report_counts"`
	UniqueCrashCounts []CrashUniqueReportCounts `json:"unique_crash_counts"`
}

type CrashReportCounts struct {
	ServerVersion string `json:"server_version"`
	CompileDate   string `json:"compile_date"`
	CrashCount    int    `json:"crash_count"`
	ResolvedCount int    `json:"resolved_count"`
}

type CrashUniqueReportCounts struct {
	ServerVersion    string `json:"server_version"`
	Fingerprint      string `json:"fingerprint"`
	UniqueCrashCount int    `json:"unique_crash_count"`
	ResolvedCount    int    `json:"resolved_count"`
}

func (a *Controller) getServerCrashReportCounts(c echo.Context) error {
	db, err := a.db.GetSpireDb().DB()
	if err != nil {
		return err
	}

	rows, err := db.Query("select server_version, compile_date, count(*) as crash_count, sum(resolved) as resolved from spire_crash_reports group by server_version order by server_version")
	if err != nil {
		return err
	}

	var counts []CrashReportCounts
	for rows.Next() {
		var r CrashReportCounts
		err = rows.Scan(&r.ServerVersion, &r.CompileDate, &r.CrashCount, &r.ResolvedCount)
		if err != nil {
			return err
		}
		counts = append(counts, r)
	}

	rows, err = db.Query("select server_version, fingerprint, count(*) as crash_count, sum(resolved) as resolved from spire_crash_reports group by server_version, fingerprint order by server_version")
	if err != nil {
		return err
	}

	var uniqueCounts []CrashUniqueReportCounts
	for rows.Next() {
		var r CrashUniqueReportCounts
		err = rows.Scan(&r.ServerVersion, &r.Fingerprint, &r.UniqueCrashCount, &r.ResolvedCount)
		if err != nil {
			return err
		}
		uniqueCounts = append(uniqueCounts, r)
	}

	return c.JSON(
		http.StatusOK,
		CrashReportResponse{
			CrashReportCounts: counts,
			UniqueCrashCounts: uniqueCounts,
		},
	)
}

func (a *Controller) getReleases(c echo.Context) error {
	r, err := a.releases.getReleases()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(
		http.StatusOK,
		echo.Map{"data": r},
	)
}
