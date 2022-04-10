package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type BugReportController struct {
	db	 *database.DatabaseResolver
	logger *logrus.Logger
}

func NewBugReportController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *BugReportController {
	return &BugReportController{
		db:	 db,
		logger: logger,
	}
}

func (e *BugReportController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "bug_report/:id", e.getBugReport, nil),
		routes.RegisterRoute(http.MethodGet, "bug_reports", e.listBugReports, nil),
		routes.RegisterRoute(http.MethodPut, "bug_report", e.createBugReport, nil),
		routes.RegisterRoute(http.MethodDelete, "bug_report/:id", e.deleteBugReport, nil),
		routes.RegisterRoute(http.MethodPatch, "bug_report/:id", e.updateBugReport, nil),
		routes.RegisterRoute(http.MethodPost, "bug_reports/bulk", e.getBugReportsBulk, nil),
	}
}

// listBugReports godoc
// @Id listBugReports
// @Summary Lists BugReports
// @Accept json
// @Produce json
// @Tags BugReport
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BugReport
// @Failure 500 {string} string "Bad query request"
// @Router /bug_reports [get]
func (e *BugReportController) listBugReports(c echo.Context) error {
	var results []models.BugReport
	err := e.db.QueryContext(models.BugReport{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getBugReport godoc
// @Id getBugReport
// @Summary Gets BugReport
// @Accept json
// @Produce json
// @Tags BugReport
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names "
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.BugReport
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /bug_report/{id} [get]
func (e *BugReportController) getBugReport(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.BugReport
	query := e.db.QueryContext(models.BugReport{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.ID == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateBugReport godoc
// @Id updateBugReport
// @Summary Updates BugReport
// @Accept json
// @Produce json
// @Tags BugReport
// @Param id path int true "Id"
// @Param bug_report body models.BugReport true "BugReport"
// @Success 200 {array} models.BugReport
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /bug_report/{id} [patch]
func (e *BugReportController) updateBugReport(c echo.Context) error {
	request := new(models.BugReport)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.BugReport
	query := e.db.QueryContext(models.BugReport{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	err = query.Select("*").Updates(&request).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	return c.JSON(http.StatusOK, request)
}

// createBugReport godoc
// @Id createBugReport
// @Summary Creates BugReport
// @Accept json
// @Produce json
// @Param bug_report body models.BugReport true "BugReport"
// @Tags BugReport
// @Success 200 {array} models.BugReport
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /bug_report [put]
func (e *BugReportController) createBugReport(c echo.Context) error {
	bugReport := new(models.BugReport)
	if err := c.Bind(bugReport); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.BugReport{}, c).Model(&models.BugReport{}).Create(&bugReport).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	return c.JSON(http.StatusOK, bugReport)
}

// deleteBugReport godoc
// @Id deleteBugReport
// @Summary Deletes BugReport
// @Accept json
// @Produce json
// @Tags BugReport
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /bug_report/{id} [delete]
func (e *BugReportController) deleteBugReport(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// query builder
	var result models.BugReport
	query := e.db.QueryContext(models.BugReport{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = e.db.Get(models.BugReport{}, c).Model(&models.BugReport{}).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getBugReportsBulk godoc
// @Id getBugReportsBulk
// @Summary Gets BugReports in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags BugReport
// @Success 200 {array} models.BugReport
// @Failure 500 {string} string "Bad query request"
// @Router /bug_reports/bulk [post]
func (e *BugReportController) getBugReportsBulk(c echo.Context) error {
	var results []models.BugReport

	r := new(BulkFetchByIdsGetRequest)
	if err := c.Bind(r); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to bulk request: [%v]", err.Error())},
		)
	}

	if len(r.IDs) == 0 {
		return c.JSON(
			http.StatusOK,
			echo.Map{"error": fmt.Sprintf("Missing request field data 'ids'")},
		)
	}

	err := e.db.QueryContext(models.BugReport{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
