package crudcontrollers

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ReportController struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
}

func NewReportController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *ReportController {
	return &ReportController{
		db:     db,
		logger: logger,
	}
}

func (e *ReportController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodDelete, "report/:report", e.deleteReport, nil),
		routes.RegisterRoute(http.MethodGet, "report/:report", e.getReport, nil),
		routes.RegisterRoute(http.MethodGet, "reports", e.listReports, nil),
		routes.RegisterRoute(http.MethodPost, "reports/bulk", e.getReportsBulk, nil),
		routes.RegisterRoute(http.MethodPatch, "report/:report", e.updateReport, nil),
		routes.RegisterRoute(http.MethodPut, "report", e.createReport, nil),
	}
}

// listReports godoc
// @Id listReports
// @Summary Lists Reports
// @Accept json
// @Produce json
// @Tags Report
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Report
// @Failure 500 {string} string "Bad query request"
// @Router /reports [get]
func (e *ReportController) listReports(c echo.Context) error {
	var results []models.Report
	err := e.db.QueryContext(models.Report{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getReport godoc
// @Id getReport
// @Summary Gets Report
// @Accept json
// @Produce json
// @Tags Report
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.Report
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /report/{id} [get]
func (e *ReportController) getReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("report"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param"})
	}

	var result models.Report
	err = e.db.QueryContext(models.Report{}, c).First(&result, reportId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateReport godoc
// @Id updateReport
// @Summary Updates Report
// @Accept json
// @Produce json
// @Tags Report
// @Param id path int true "Id"
// @Param report body models.Report true "Report"
// @Success 200 {array} models.Report
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /report/{id} [patch]
func (e *ReportController) updateReport(c echo.Context) error {
	report := new(models.Report)
	if err := c.Bind(report); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

    entity := models.Report{}
	err := e.db.Get(models.Report{}, c).Model(&models.Report{}).First(&entity, report.ID).Error
	if err != nil || report.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Report{}, c).Model(&entity).Updates(&report).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity: [%v]", err)})
	}

	return c.JSON(http.StatusOK, report)
}

// createReport godoc
// @Id createReport
// @Summary Creates Report
// @Accept json
// @Produce json
// @Param report body models.Report true "Report"
// @Tags Report
// @Success 200 {array} models.Report
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /report [put]
func (e *ReportController) createReport(c echo.Context) error {
	report := new(models.Report)
	if err := c.Bind(report); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity: [%v]", err)},
		)
	}

	err := e.db.Get(models.Report{}, c).Model(&models.Report{}).Create(&report).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity: [%v]", err)},
		)
	}

	return c.JSON(http.StatusOK, report)
}

// deleteReport godoc
// @Id deleteReport
// @Summary Deletes Report
// @Accept json
// @Produce json
// @Tags Report
// @Param id path int true "Id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /report/{id} [delete]
func (e *ReportController) deleteReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("report"))
	if err != nil {
		e.logger.Error(err)
	}

	report := new(models.Report)
	err = e.db.Get(models.Report{}, c).Model(&models.Report{}).First(&report, reportId).Error
	if err != nil || report.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	err = e.db.Get(models.Report{}, c).Model(&models.Report{}).Delete(&report).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getReportsBulk godoc
// @Id getReportsBulk
// @Summary Gets Reports in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags Report
// @Success 200 {array} models.Report
// @Failure 500 {string} string "Bad query request"
// @Router /reports/bulk [post]
func (e *ReportController) getReportsBulk(c echo.Context) error {
	var results []models.Report

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err)},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.Report{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}
