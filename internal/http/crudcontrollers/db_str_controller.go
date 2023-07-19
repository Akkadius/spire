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
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"
)

type DbStrController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewDbStrController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *DbStrController {
	return &DbStrController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *DbStrController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "db_str/:id", e.getDbStr, nil),
		routes.RegisterRoute(http.MethodGet, "db_strs", e.listDbStrs, nil),
		routes.RegisterRoute(http.MethodGet, "db_strs/count", e.getDbStrsCount, nil),
		routes.RegisterRoute(http.MethodPut, "db_str", e.createDbStr, nil),
		routes.RegisterRoute(http.MethodDelete, "db_str/:id", e.deleteDbStr, nil),
		routes.RegisterRoute(http.MethodPatch, "db_str/:id", e.updateDbStr, nil),
		routes.RegisterRoute(http.MethodPost, "db_strs/bulk", e.getDbStrsBulk, nil),
	}
}

// listDbStrs godoc
// @Id listDbStrs
// @Summary Lists DbStrs
// @Accept json
// @Produce json
// @Tags DbStr
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Bad query request"
// @Router /db_strs [get]
func (e *DbStrController) listDbStrs(c echo.Context) error {
	var results []models.DbStr
	err := e.db.QueryContext(models.DbStr{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getDbStr godoc
// @Id getDbStr
// @Summary Gets DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DbStr
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /db_str/{id} [get]
func (e *DbStrController) getDbStr(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [Id]"})
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [typeId] position [2] type [int]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.DbStr
	query := e.db.QueryContext(models.DbStr{}, c)
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

// updateDbStr godoc
// @Id updateDbStr
// @Summary Updates DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "Id"
// @Param db_str body models.DbStr true "DbStr"
// @Success 200 {array} models.DbStr
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /db_str/{id} [patch]
func (e *DbStrController) updateDbStr(c echo.Context) error {
	request := new(models.DbStr)
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

	// key param [typeId] position [2] type [int]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.DbStr
	query := e.db.QueryContext(models.DbStr{}, c)
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
		event := fmt.Sprintf("Updated [DbStr] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createDbStr godoc
// @Id createDbStr
// @Summary Creates DbStr
// @Accept json
// @Produce json
// @Param db_str body models.DbStr true "DbStr"
// @Tags DbStr
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /db_str [put]
func (e *DbStrController) createDbStr(c echo.Context) error {
	dbStr := new(models.DbStr)
	if err := c.Bind(dbStr); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.DbStr{}, c).Model(&models.DbStr{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&dbStr).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.DbStr{}, dbStr)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [DbStr] [%v] data [%v]", dbStr.ID, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, dbStr)
}

// deleteDbStr godoc
// @Id deleteDbStr
// @Summary Deletes DbStr
// @Accept json
// @Produce json
// @Tags DbStr
// @Param id path int true "id"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /db_str/{id} [delete]
func (e *DbStrController) deleteDbStr(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, id)
	keys = append(keys, "id = ?")

	// key param [typeId] position [2] type [int]
	if len(c.QueryParam("typeId")) > 0 {
		typeIdParam, err := strconv.Atoi(c.QueryParam("typeId"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [typeId] err [%s]", err.Error())})
		}

		params = append(params, typeIdParam)
		keys = append(keys, "typeId = ?")
	}

	// query builder
	var result models.DbStr
	query := e.db.QueryContext(models.DbStr{}, c)
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
		event := fmt.Sprintf("Deleted [DbStr] [%v] keys [%v]", result.ID, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getDbStrsBulk godoc
// @Id getDbStrsBulk
// @Summary Gets DbStrs in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags DbStr
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Bad query request"
// @Router /db_strs/bulk [post]
func (e *DbStrController) getDbStrsBulk(c echo.Context) error {
	var results []models.DbStr

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

	err := e.db.QueryContext(models.DbStr{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getDbStrsCount godoc
// @Id getDbStrsCount
// @Summary Counts DbStrs
// @Accept json
// @Produce json
// @Tags DbStr
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.DbStr
// @Failure 500 {string} string "Bad query request"
// @Router /db_strs/count [get]
func (e *DbStrController) getDbStrsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.DbStr{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}