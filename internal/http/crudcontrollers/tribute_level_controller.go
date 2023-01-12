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

type TributeLevelController struct {
	db       *database.DatabaseResolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewTributeLevelController(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *TributeLevelController {
	return &TributeLevelController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *TributeLevelController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "tribute_level/:tributeId", e.getTributeLevel, nil),
		routes.RegisterRoute(http.MethodGet, "tribute_levels", e.listTributeLevels, nil),
		routes.RegisterRoute(http.MethodGet, "tribute_levels/count", e.getTributeLevelsCount, nil),
		routes.RegisterRoute(http.MethodPut, "tribute_level", e.createTributeLevel, nil),
		routes.RegisterRoute(http.MethodDelete, "tribute_level/:tributeId", e.deleteTributeLevel, nil),
		routes.RegisterRoute(http.MethodPatch, "tribute_level/:tributeId", e.updateTributeLevel, nil),
		routes.RegisterRoute(http.MethodPost, "tribute_levels/bulk", e.getTributeLevelsBulk, nil),
	}
}

// listTributeLevels godoc
// @Id listTributeLevels
// @Summary Lists TributeLevels
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_levels [get]
func (e *TributeLevelController) listTributeLevels(c echo.Context) error {
	var results []models.TributeLevel
	err := e.db.QueryContext(models.TributeLevel{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, results)
}

// getTributeLevel godoc
// @Id getTributeLevel
// @Summary Gets TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TributeLevel
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_level/{id} [get]
func (e *TributeLevelController) getTributeLevel(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TributeId]"})
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.TributeId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateTributeLevel godoc
// @Id updateTributeLevel
// @Summary Updates TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "Id"
// @Param tribute_level body models.TributeLevel true "TributeLevel"
// @Success 200 {array} models.TributeLevel
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /tribute_level/{id} [patch]
func (e *TributeLevelController) updateTributeLevel(c echo.Context) error {
	request := new(models.TributeLevel)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [TributeId]"})
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
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
		event := fmt.Sprintf("Updated [TributeLevel] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createTributeLevel godoc
// @Id createTributeLevel
// @Summary Creates TributeLevel
// @Accept json
// @Produce json
// @Param tribute_level body models.TributeLevel true "TributeLevel"
// @Tags TributeLevel
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /tribute_level [put]
func (e *TributeLevelController) createTributeLevel(c echo.Context) error {
	tributeLevel := new(models.TributeLevel)
	if err := c.Bind(tributeLevel); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	err := e.db.Get(models.TributeLevel{}, c).Model(&models.TributeLevel{}).Create(&tributeLevel).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.TributeLevel{}, tributeLevel)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [TributeLevel] [%v] data [%v]", tributeLevel.TributeId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, tributeLevel)
}

// deleteTributeLevel godoc
// @Id deleteTributeLevel
// @Summary Deletes TributeLevel
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param id path int true "tributeId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /tribute_level/{id} [delete]
func (e *TributeLevelController) deleteTributeLevel(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	tributeId, err := strconv.Atoi(c.Param("tributeId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, tributeId)
	keys = append(keys, "tribute_id = ?")

	// key param [level] position [2] type [int]
	if len(c.QueryParam("level")) > 0 {
		levelParam, err := strconv.Atoi(c.QueryParam("level"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [level] err [%s]", err.Error())})
		}

		params = append(params, levelParam)
		keys = append(keys, "level = ?")
	}

	// query builder
	var result models.TributeLevel
	query := e.db.QueryContext(models.TributeLevel{}, c)
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
		event := fmt.Sprintf("Deleted [TributeLevel] [%v] keys [%v]", result.TributeId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getTributeLevelsBulk godoc
// @Id getTributeLevelsBulk
// @Summary Gets TributeLevels in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags TributeLevel
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_levels/bulk [post]
func (e *TributeLevelController) getTributeLevelsBulk(c echo.Context) error {
	var results []models.TributeLevel

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

	err := e.db.QueryContext(models.TributeLevel{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getTributeLevelsCount godoc
// @Id getTributeLevelsCount
// @Summary Counts TributeLevels
// @Accept json
// @Produce json
// @Tags TributeLevel
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.TributeLevel
// @Failure 500 {string} string "Bad query request"
// @Router /tribute_levels/count [get]
func (e *TributeLevelController) getTributeLevelsCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.TributeLevel{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}