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

type FactionValueController struct {
	db       *database.Resolver
	logger   *logrus.Logger
	auditLog *auditlog.UserEvent
}

func NewFactionValueController(
	db *database.Resolver,
	logger *logrus.Logger,
	auditLog *auditlog.UserEvent,
) *FactionValueController {
	return &FactionValueController{
		db:       db,
		logger:   logger,
		auditLog: auditLog,
	}
}

func (e *FactionValueController) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "faction_value/:charId", e.getFactionValue, nil),
		routes.RegisterRoute(http.MethodGet, "faction_values", e.listFactionValues, nil),
		routes.RegisterRoute(http.MethodGet, "faction_values/count", e.getFactionValuesCount, nil),
		routes.RegisterRoute(http.MethodPut, "faction_value", e.createFactionValue, nil),
		routes.RegisterRoute(http.MethodDelete, "faction_value/:charId", e.deleteFactionValue, nil),
		routes.RegisterRoute(http.MethodPatch, "faction_value/:charId", e.updateFactionValue, nil),
		routes.RegisterRoute(http.MethodPost, "faction_values/bulk", e.getFactionValuesBulk, nil),
	}
}

// listFactionValues godoc
// @Id listFactionValues
// @Summary Lists FactionValues
// @Accept json
// @Produce json
// @Tags FactionValue
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionValue
// @Failure 500 {string} string "Bad query request"
// @Router /faction_values [get]
func (e *FactionValueController) listFactionValues(c echo.Context) error {
	var results []models.FactionValue
	err := e.db.QueryContext(models.FactionValue{}, c).Find(&results).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getFactionValue godoc
// @Id getFactionValue
// @Summary Gets FactionValue
// @Accept json
// @Produce json
// @Tags FactionValue
// @Param id path int true "Id"
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionValue
// @Failure 404 {string} string "Entity not found"
// @Failure 500 {string} string "Cannot find param"
// @Failure 500 {string} string "Bad query request"
// @Router /faction_value/{id} [get]
func (e *FactionValueController) getFactionValue(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.FactionValue
	query := e.db.QueryContext(models.FactionValue{}, c)
	for i, _ := range keys {
		query = query.Where(keys[i], params[i])
	}

	// grab first entry
	err = query.First(&result).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// couldn't find entity
	if result.CharId == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Cannot find entity"})
	}

	return c.JSON(http.StatusOK, result)
}

// updateFactionValue godoc
// @Id updateFactionValue
// @Summary Updates FactionValue
// @Accept json
// @Produce json
// @Tags FactionValue
// @Param id path int true "Id"
// @Param faction_value body models.FactionValue true "FactionValue"
// @Success 200 {array} models.FactionValue
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error updating entity"
// @Router /faction_value/{id} [patch]
func (e *FactionValueController) updateFactionValue(c echo.Context) error {
	request := new(models.FactionValue)
	if err := c.Bind(request); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Cannot find param [CharId]"})
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.FactionValue
	query := e.db.QueryContext(models.FactionValue{}, c)
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
		event := fmt.Sprintf("Updated [FactionValue] [%v] fields [%v]", strings.Join(ids, ", "), strings.Join(fieldsUpdated, ", "))
		e.auditLog.LogUserEvent(c, "UPDATE", event)
	}

	return c.JSON(http.StatusOK, request)
}

// createFactionValue godoc
// @Id createFactionValue
// @Summary Creates FactionValue
// @Accept json
// @Produce json
// @Param faction_value body models.FactionValue true "FactionValue"
// @Tags FactionValue
// @Success 200 {array} models.FactionValue
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error inserting entity"
// @Router /faction_value [put]
func (e *FactionValueController) createFactionValue(c echo.Context) error {
	factionValue := new(models.FactionValue)
	if err := c.Bind(factionValue); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error binding to entity [%v]", err.Error())},
		)
	}

	db := e.db.Get(models.FactionValue{}, c).Model(&models.FactionValue{})

	// save associations
	if c.QueryParam("save_associations") != "true" {
		db = db.Omit(clause.Associations)
	}

	err := db.Create(&factionValue).Error
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"error": fmt.Sprintf("Error inserting entity [%v]", err.Error())},
		)
	}

	// log create event
	if e.db.GetSpireDb() != nil {
		// diff between an empty model and the created
		diff := database.ResultDifference(models.FactionValue{}, factionValue)
		// build fields created
		var fields []string
		for k, v := range diff {
			fields = append(fields, fmt.Sprintf("%v = %v", k, v))
		}
		// record event
		event := fmt.Sprintf("Created [FactionValue] [%v] data [%v]", factionValue.CharId, strings.Join(fields, ", "))
		e.auditLog.LogUserEvent(c, "CREATE", event)
	}

	return c.JSON(http.StatusOK, factionValue)
}

// deleteFactionValue godoc
// @Id deleteFactionValue
// @Summary Deletes FactionValue
// @Accept json
// @Produce json
// @Tags FactionValue
// @Param id path int true "charId"
// @Success 200 {string} string "Entity deleted successfully"
// @Failure 404 {string} string "Cannot find entity"
// @Failure 500 {string} string "Error binding to entity"
// @Failure 500 {string} string "Error deleting entity"
// @Router /faction_value/{id} [delete]
func (e *FactionValueController) deleteFactionValue(c echo.Context) error {
	var params []interface{}
	var keys []string

	// primary key param
	charId, err := strconv.Atoi(c.Param("charId"))
	if err != nil {
		e.logger.Error(err)
	}
	params = append(params, charId)
	keys = append(keys, "char_id = ?")

	// key param [faction_id] position [2] type [int]
	if len(c.QueryParam("faction_id")) > 0 {
		factionIdParam, err := strconv.Atoi(c.QueryParam("faction_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": fmt.Sprintf("Error parsing query param [faction_id] err [%s]", err.Error())})
		}

		params = append(params, factionIdParam)
		keys = append(keys, "faction_id = ?")
	}

	// query builder
	var result models.FactionValue
	query := e.db.QueryContext(models.FactionValue{}, c)
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
		event := fmt.Sprintf("Deleted [FactionValue] [%v] keys [%v]", result.CharId, strings.Join(ids, ", "))
		e.auditLog.LogUserEvent(c, "DELETE", event)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": "Entity deleted successfully"})
}

// getFactionValuesBulk godoc
// @Id getFactionValuesBulk
// @Summary Gets FactionValues in bulk
// @Accept json
// @Produce json
// @Param Body body BulkFetchByIdsGetRequest true "body"
// @Tags FactionValue
// @Success 200 {array} models.FactionValue
// @Failure 500 {string} string "Bad query request"
// @Router /faction_values/bulk [post]
func (e *FactionValueController) getFactionValuesBulk(c echo.Context) error {
	var results []models.FactionValue

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

	err := e.db.QueryContext(models.FactionValue{}, c).Find(&results, r.IDs).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, results)
}

// getFactionValuesCount godoc
// @Id getFactionValuesCount
// @Summary Counts FactionValues
// @Accept json
// @Produce json
// @Tags FactionValue
// @Param includes query string false "Relationships [all] for all [number] for depth of relationships to load or [.] separated relationship names"
// @Param where query string false "Filter on specific fields. Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param whereOr query string false "Filter on specific fields (Chained ors). Multiple conditions [.] separated Example: col_like_value.col2__val2"
// @Param groupBy query string false "Group by field. Multiple conditions [.] separated Example: field1.field2"
// @Param limit query string false "Rows to limit in response (Default: 10,000)"
// @Param page query int 0 "Pagination page"
// @Param orderBy query string false "Order by [field]"
// @Param orderDirection query string false "Order by field direction"
// @Param select query string false "Column names [.] separated to fetch specific fields in response"
// @Success 200 {array} models.FactionValue
// @Failure 500 {string} string "Bad query request"
// @Router /faction_values/count [get]
func (e *FactionValueController) getFactionValuesCount(c echo.Context) error {
	var count int64
	err := e.db.QueryContext(models.FactionValue{}, c).Count(&count).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"count": count})
}