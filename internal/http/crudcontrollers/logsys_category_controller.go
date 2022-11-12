package crudcontrollers

import (
	"fmt"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type LogsysCategoryController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewLogsysCategoryController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *LogsysCategoryController {
	return &LogsysCategoryController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *LogsysCategoryController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "logsys_category/:logCategoryId", e.getLogsysCategory, nil),
		routes.RegisterRoute(http.MethodGet, "logsys_categories", e.listLogsysCategories, nil),
		routes.RegisterRoute(http.MethodPut, "logsys_category", e.createLogsysCategory, nil),
		routes.RegisterRoute(http.MethodDelete, "logsys_category/:logCategoryId", e.deleteLogsysCategory, nil),
		routes.RegisterRoute(http.MethodPatch, "logsys_category/:logCategoryId", e.updateLogsysCategory, nil),
		routes.RegisterRoute(http.MethodPost, "logsys_categories/bulk", e.getLogsysCategoriesBulk, nil),
	}
}

// listLogsysCategories godoc
// @Id listLogsysCategories
// @Summary Lists LogsysCategories
// @Accept json
// @Produce json
// @Tags LogsysCategory
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LogsysCategory
// @Failure 500 {string} string "Bad query request"
// @Router /logsys_categories [get]
func (e *LogsysCategoryController) listLogsysCategories(c echo.Context) error {
	var results []models.LogsysCategory
	err := e.db.QueryContext(models.LogsysCategory{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getLogsysCategory godoc
// @Id getLogsysCategory
// @Summary Gets LogsysCategory
// @Accept json
// @Produce json
// @Tags LogsysCategory
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.LogsysCategory
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /logsys_category/{id} [get]
func (e *LogsysCategoryController) getLogsysCategory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	logCategoryId, err := strconv.Atoi(c.Param("logCategoryId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LogCategoryId]"})
	}
	params = append(params, logCategoryId)
	keys = append(keys, "log_category_id = ?")

	// query builder
	var result models.LogsysCategory
	query := e.db.QueryContext(models.LogsysCategory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.LogCategoryId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateLogsysCategory godoc
// @Id updateLogsysCategory
// @Summary Updates LogsysCategory
// @Accept json
// @Produce json
// @Tags LogsysCategory
// @Param id path int true "Id"
// @Param logsys_category body models.LogsysCategory true "LogsysCategory"
// @Success 200 {array} models.LogsysCategory
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /logsys_category/{id} [patch]
func (e *LogsysCategoryController) updateLogsysCategory(c echo.Context) error {
	request := new(models.LogsysCategory)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	logCategoryId, err := strconv.Atoi(c.Param("logCategoryId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [LogCategoryId]"})
	}
	params = append(params, logCategoryId)
	keys = append(keys, "log_category_id = ?")

	// query builder
	var result models.LogsysCategory
	query := e.db.QueryContext(models.LogsysCategory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Cannot find entity [%s]", err.Error())})
	}

	// save top-level using only changes
	diff := database.ResultDifference(result, request)
	err = query.Session(&gorm.Session{FullSaveAssociations: false}).Updates(diff).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error updating entity [%v]", err.Error())})
	}

	// log update event
	if e.db.GetSpireDb() != nil && len(diff) > 0 {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// build fields updated
		var fieldsUpdated []string
		for k, v := range diff {
			fieldsUpdated = append(fieldsUpdated, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Updated [LogsysCategory] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createLogsysCategory godoc
// @Id createLogsysCategory
// @Summary Creates LogsysCategory
// @Accept json
// @Produce json
// @Param logsys_category body models.LogsysCategory true "LogsysCategory"
// @Tags LogsysCategory
// @Success 200 {array} models.LogsysCategory
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /logsys_category [put]
func (e *LogsysCategoryController) createLogsysCategory(c echo.Context) error {
	logsysCategory := new(models.LogsysCategory)
	if err := c.Bind(logsysCategory); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.LogsysCategory{}, c).Model(&models.LogsysCategory{}).Create(&logsysCategory).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.LogsysCategory{}, logsysCategory)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [LogsysCategory] [%v] data [%v]", logsysCategory.LogCategoryId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, logsysCategory)
}

// deleteLogsysCategory godoc
// @Id deleteLogsysCategory
// @Summary Deletes LogsysCategory
// @Accept json
// @Produce json
// @Tags LogsysCategory
// @Param id path int true "logCategoryId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /logsys_category/{id} [delete]
func (e *LogsysCategoryController) deleteLogsysCategory(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	logCategoryId, err := strconv.Atoi(c.Param("logCategoryId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, logCategoryId)
	keys = append(keys, "log_category_id = ?")

	// query builder
	var result models.LogsysCategory
	query := e.db.QueryContext(models.LogsysCategory{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	err = query.Limit(10000).Delete(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error deleting entity"})
	}

	// log delete event
	if e.db.GetSpireDb() != nil {
		// build ids
		var ids []string
		for i, _ := range keys {
			param := fmt.Sprintf("%v", params[i])
			ids = append(ids, fmt.Sprintf("%v", strings.ReplaceAll(keys[i], "?", param)))
		}
		// record event
		event := fmt.Sprintf("Deleted [LogsysCategory] [%v] keys [%v]", result.LogCategoryId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getLogsysCategoriesBulk godoc
// @Id getLogsysCategoriesBulk
// @Summary Gets LogsysCategories in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags LogsysCategory
// @Success 200 {array} models.LogsysCategory
// @Failure 500 {string} string "Bad query request"
// @Router /logsys_categories/bulk [post]
func (e *LogsysCategoryController) getLogsysCategoriesBulk(c echo.Context) error {
	var results []models.LogsysCategory

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

	err := e.db.QueryContext(models.LogsysCategory{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}
